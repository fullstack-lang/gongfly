// generated code - do not edit
package models

import (
	"errors"
	"fmt"
	"math"
	"time"
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
	DummyAgents           map[*DummyAgent]any
	DummyAgents_mapString map[string]*DummyAgent

	// insertion point for slice of pointers maps

	OnAfterDummyAgentCreateCallback OnAfterCreateInterface[DummyAgent]
	OnAfterDummyAgentUpdateCallback OnAfterUpdateInterface[DummyAgent]
	OnAfterDummyAgentDeleteCallback OnAfterDeleteInterface[DummyAgent]
	OnAfterDummyAgentReadCallback   OnAfterReadInterface[DummyAgent]

	Engines           map[*Engine]any
	Engines_mapString map[string]*Engine

	// insertion point for slice of pointers maps

	OnAfterEngineCreateCallback OnAfterCreateInterface[Engine]
	OnAfterEngineUpdateCallback OnAfterUpdateInterface[Engine]
	OnAfterEngineDeleteCallback OnAfterDeleteInterface[Engine]
	OnAfterEngineReadCallback   OnAfterReadInterface[Engine]

	Events           map[*Event]any
	Events_mapString map[string]*Event

	// insertion point for slice of pointers maps

	OnAfterEventCreateCallback OnAfterCreateInterface[Event]
	OnAfterEventUpdateCallback OnAfterUpdateInterface[Event]
	OnAfterEventDeleteCallback OnAfterDeleteInterface[Event]
	OnAfterEventReadCallback   OnAfterReadInterface[Event]

	GongsimCommands           map[*GongsimCommand]any
	GongsimCommands_mapString map[string]*GongsimCommand

	// insertion point for slice of pointers maps

	OnAfterGongsimCommandCreateCallback OnAfterCreateInterface[GongsimCommand]
	OnAfterGongsimCommandUpdateCallback OnAfterUpdateInterface[GongsimCommand]
	OnAfterGongsimCommandDeleteCallback OnAfterDeleteInterface[GongsimCommand]
	OnAfterGongsimCommandReadCallback   OnAfterReadInterface[GongsimCommand]

	GongsimStatuss           map[*GongsimStatus]any
	GongsimStatuss_mapString map[string]*GongsimStatus

	// insertion point for slice of pointers maps

	OnAfterGongsimStatusCreateCallback OnAfterCreateInterface[GongsimStatus]
	OnAfterGongsimStatusUpdateCallback OnAfterUpdateInterface[GongsimStatus]
	OnAfterGongsimStatusDeleteCallback OnAfterDeleteInterface[GongsimStatus]
	OnAfterGongsimStatusReadCallback   OnAfterReadInterface[GongsimStatus]

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
	return "github.com/fullstack-lang/gongsim/go/models"
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
	CommitDummyAgent(dummyagent *DummyAgent)
	CheckoutDummyAgent(dummyagent *DummyAgent)
	CommitEngine(engine *Engine)
	CheckoutEngine(engine *Engine)
	CommitEvent(event *Event)
	CheckoutEvent(event *Event)
	CommitGongsimCommand(gongsimcommand *GongsimCommand)
	CheckoutGongsimCommand(gongsimcommand *GongsimCommand)
	CommitGongsimStatus(gongsimstatus *GongsimStatus)
	CheckoutGongsimStatus(gongsimstatus *GongsimStatus)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		DummyAgents:           make(map[*DummyAgent]any),
		DummyAgents_mapString: make(map[string]*DummyAgent),

		Engines:           make(map[*Engine]any),
		Engines_mapString: make(map[string]*Engine),

		Events:           make(map[*Event]any),
		Events_mapString: make(map[string]*Event),

		GongsimCommands:           make(map[*GongsimCommand]any),
		GongsimCommands_mapString: make(map[string]*GongsimCommand),

		GongsimStatuss:           make(map[*GongsimStatus]any),
		GongsimStatuss_mapString: make(map[string]*GongsimStatus),

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
	stage.Map_GongStructName_InstancesNb["DummyAgent"] = len(stage.DummyAgents)
	stage.Map_GongStructName_InstancesNb["Engine"] = len(stage.Engines)
	stage.Map_GongStructName_InstancesNb["Event"] = len(stage.Events)
	stage.Map_GongStructName_InstancesNb["GongsimCommand"] = len(stage.GongsimCommands)
	stage.Map_GongStructName_InstancesNb["GongsimStatus"] = len(stage.GongsimStatuss)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["DummyAgent"] = len(stage.DummyAgents)
	stage.Map_GongStructName_InstancesNb["Engine"] = len(stage.Engines)
	stage.Map_GongStructName_InstancesNb["Event"] = len(stage.Events)
	stage.Map_GongStructName_InstancesNb["GongsimCommand"] = len(stage.GongsimCommands)
	stage.Map_GongStructName_InstancesNb["GongsimStatus"] = len(stage.GongsimStatuss)

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
// Stage puts dummyagent to the model stage
func (dummyagent *DummyAgent) Stage(stage *StageStruct) *DummyAgent {
	stage.DummyAgents[dummyagent] = __member
	stage.DummyAgents_mapString[dummyagent.Name] = dummyagent

	return dummyagent
}

// Unstage removes dummyagent off the model stage
func (dummyagent *DummyAgent) Unstage(stage *StageStruct) *DummyAgent {
	delete(stage.DummyAgents, dummyagent)
	delete(stage.DummyAgents_mapString, dummyagent.Name)
	return dummyagent
}

// UnstageVoid removes dummyagent off the model stage
func (dummyagent *DummyAgent) UnstageVoid(stage *StageStruct) {
	delete(stage.DummyAgents, dummyagent)
	delete(stage.DummyAgents_mapString, dummyagent.Name)
}

// commit dummyagent to the back repo (if it is already staged)
func (dummyagent *DummyAgent) Commit(stage *StageStruct) *DummyAgent {
	if _, ok := stage.DummyAgents[dummyagent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDummyAgent(dummyagent)
		}
	}
	return dummyagent
}

func (dummyagent *DummyAgent) CommitVoid(stage *StageStruct) {
	dummyagent.Commit(stage)
}

// Checkout dummyagent to the back repo (if it is already staged)
func (dummyagent *DummyAgent) Checkout(stage *StageStruct) *DummyAgent {
	if _, ok := stage.DummyAgents[dummyagent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDummyAgent(dummyagent)
		}
	}
	return dummyagent
}

// for satisfaction of GongStruct interface
func (dummyagent *DummyAgent) GetName() (res string) {
	return dummyagent.Name
}

// Stage puts engine to the model stage
func (engine *Engine) Stage(stage *StageStruct) *Engine {
	stage.Engines[engine] = __member
	stage.Engines_mapString[engine.Name] = engine

	return engine
}

// Unstage removes engine off the model stage
func (engine *Engine) Unstage(stage *StageStruct) *Engine {
	delete(stage.Engines, engine)
	delete(stage.Engines_mapString, engine.Name)
	return engine
}

// UnstageVoid removes engine off the model stage
func (engine *Engine) UnstageVoid(stage *StageStruct) {
	delete(stage.Engines, engine)
	delete(stage.Engines_mapString, engine.Name)
}

// commit engine to the back repo (if it is already staged)
func (engine *Engine) Commit(stage *StageStruct) *Engine {
	if _, ok := stage.Engines[engine]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEngine(engine)
		}
	}
	return engine
}

func (engine *Engine) CommitVoid(stage *StageStruct) {
	engine.Commit(stage)
}

// Checkout engine to the back repo (if it is already staged)
func (engine *Engine) Checkout(stage *StageStruct) *Engine {
	if _, ok := stage.Engines[engine]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEngine(engine)
		}
	}
	return engine
}

// for satisfaction of GongStruct interface
func (engine *Engine) GetName() (res string) {
	return engine.Name
}

// Stage puts event to the model stage
func (event *Event) Stage(stage *StageStruct) *Event {
	stage.Events[event] = __member
	stage.Events_mapString[event.Name] = event

	return event
}

// Unstage removes event off the model stage
func (event *Event) Unstage(stage *StageStruct) *Event {
	delete(stage.Events, event)
	delete(stage.Events_mapString, event.Name)
	return event
}

// UnstageVoid removes event off the model stage
func (event *Event) UnstageVoid(stage *StageStruct) {
	delete(stage.Events, event)
	delete(stage.Events_mapString, event.Name)
}

// commit event to the back repo (if it is already staged)
func (event *Event) Commit(stage *StageStruct) *Event {
	if _, ok := stage.Events[event]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEvent(event)
		}
	}
	return event
}

func (event *Event) CommitVoid(stage *StageStruct) {
	event.Commit(stage)
}

// Checkout event to the back repo (if it is already staged)
func (event *Event) Checkout(stage *StageStruct) *Event {
	if _, ok := stage.Events[event]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEvent(event)
		}
	}
	return event
}

// for satisfaction of GongStruct interface
func (event *Event) GetName() (res string) {
	return event.Name
}

// Stage puts gongsimcommand to the model stage
func (gongsimcommand *GongsimCommand) Stage(stage *StageStruct) *GongsimCommand {
	stage.GongsimCommands[gongsimcommand] = __member
	stage.GongsimCommands_mapString[gongsimcommand.Name] = gongsimcommand

	return gongsimcommand
}

// Unstage removes gongsimcommand off the model stage
func (gongsimcommand *GongsimCommand) Unstage(stage *StageStruct) *GongsimCommand {
	delete(stage.GongsimCommands, gongsimcommand)
	delete(stage.GongsimCommands_mapString, gongsimcommand.Name)
	return gongsimcommand
}

// UnstageVoid removes gongsimcommand off the model stage
func (gongsimcommand *GongsimCommand) UnstageVoid(stage *StageStruct) {
	delete(stage.GongsimCommands, gongsimcommand)
	delete(stage.GongsimCommands_mapString, gongsimcommand.Name)
}

// commit gongsimcommand to the back repo (if it is already staged)
func (gongsimcommand *GongsimCommand) Commit(stage *StageStruct) *GongsimCommand {
	if _, ok := stage.GongsimCommands[gongsimcommand]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongsimCommand(gongsimcommand)
		}
	}
	return gongsimcommand
}

func (gongsimcommand *GongsimCommand) CommitVoid(stage *StageStruct) {
	gongsimcommand.Commit(stage)
}

// Checkout gongsimcommand to the back repo (if it is already staged)
func (gongsimcommand *GongsimCommand) Checkout(stage *StageStruct) *GongsimCommand {
	if _, ok := stage.GongsimCommands[gongsimcommand]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongsimCommand(gongsimcommand)
		}
	}
	return gongsimcommand
}

// for satisfaction of GongStruct interface
func (gongsimcommand *GongsimCommand) GetName() (res string) {
	return gongsimcommand.Name
}

// Stage puts gongsimstatus to the model stage
func (gongsimstatus *GongsimStatus) Stage(stage *StageStruct) *GongsimStatus {
	stage.GongsimStatuss[gongsimstatus] = __member
	stage.GongsimStatuss_mapString[gongsimstatus.Name] = gongsimstatus

	return gongsimstatus
}

// Unstage removes gongsimstatus off the model stage
func (gongsimstatus *GongsimStatus) Unstage(stage *StageStruct) *GongsimStatus {
	delete(stage.GongsimStatuss, gongsimstatus)
	delete(stage.GongsimStatuss_mapString, gongsimstatus.Name)
	return gongsimstatus
}

// UnstageVoid removes gongsimstatus off the model stage
func (gongsimstatus *GongsimStatus) UnstageVoid(stage *StageStruct) {
	delete(stage.GongsimStatuss, gongsimstatus)
	delete(stage.GongsimStatuss_mapString, gongsimstatus.Name)
}

// commit gongsimstatus to the back repo (if it is already staged)
func (gongsimstatus *GongsimStatus) Commit(stage *StageStruct) *GongsimStatus {
	if _, ok := stage.GongsimStatuss[gongsimstatus]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongsimStatus(gongsimstatus)
		}
	}
	return gongsimstatus
}

func (gongsimstatus *GongsimStatus) CommitVoid(stage *StageStruct) {
	gongsimstatus.Commit(stage)
}

// Checkout gongsimstatus to the back repo (if it is already staged)
func (gongsimstatus *GongsimStatus) Checkout(stage *StageStruct) *GongsimStatus {
	if _, ok := stage.GongsimStatuss[gongsimstatus]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongsimStatus(gongsimstatus)
		}
	}
	return gongsimstatus
}

// for satisfaction of GongStruct interface
func (gongsimstatus *GongsimStatus) GetName() (res string) {
	return gongsimstatus.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMDummyAgent(DummyAgent *DummyAgent)
	CreateORMEngine(Engine *Engine)
	CreateORMEvent(Event *Event)
	CreateORMGongsimCommand(GongsimCommand *GongsimCommand)
	CreateORMGongsimStatus(GongsimStatus *GongsimStatus)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMDummyAgent(DummyAgent *DummyAgent)
	DeleteORMEngine(Engine *Engine)
	DeleteORMEvent(Event *Event)
	DeleteORMGongsimCommand(GongsimCommand *GongsimCommand)
	DeleteORMGongsimStatus(GongsimStatus *GongsimStatus)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.DummyAgents = make(map[*DummyAgent]any)
	stage.DummyAgents_mapString = make(map[string]*DummyAgent)

	stage.Engines = make(map[*Engine]any)
	stage.Engines_mapString = make(map[string]*Engine)

	stage.Events = make(map[*Event]any)
	stage.Events_mapString = make(map[string]*Event)

	stage.GongsimCommands = make(map[*GongsimCommand]any)
	stage.GongsimCommands_mapString = make(map[string]*GongsimCommand)

	stage.GongsimStatuss = make(map[*GongsimStatus]any)
	stage.GongsimStatuss_mapString = make(map[string]*GongsimStatus)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.DummyAgents = nil
	stage.DummyAgents_mapString = nil

	stage.Engines = nil
	stage.Engines_mapString = nil

	stage.Events = nil
	stage.Events_mapString = nil

	stage.GongsimCommands = nil
	stage.GongsimCommands_mapString = nil

	stage.GongsimStatuss = nil
	stage.GongsimStatuss_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for dummyagent := range stage.DummyAgents {
		dummyagent.Unstage(stage)
	}

	for engine := range stage.Engines {
		engine.Unstage(stage)
	}

	for event := range stage.Events {
		event.Unstage(stage)
	}

	for gongsimcommand := range stage.GongsimCommands {
		gongsimcommand.Unstage(stage)
	}

	for gongsimstatus := range stage.GongsimStatuss {
		gongsimstatus.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	DummyAgent | Engine | Event | GongsimCommand | GongsimStatus
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*DummyAgent | *Engine | *Event | *GongsimCommand | *GongsimStatus
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*DummyAgent]any |
		map[*Engine]any |
		map[*Event]any |
		map[*GongsimCommand]any |
		map[*GongsimStatus]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*DummyAgent |
		map[string]*Engine |
		map[string]*Event |
		map[string]*GongsimCommand |
		map[string]*GongsimStatus |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*DummyAgent]any:
		return any(&stage.DummyAgents).(*Type)
	case map[*Engine]any:
		return any(&stage.Engines).(*Type)
	case map[*Event]any:
		return any(&stage.Events).(*Type)
	case map[*GongsimCommand]any:
		return any(&stage.GongsimCommands).(*Type)
	case map[*GongsimStatus]any:
		return any(&stage.GongsimStatuss).(*Type)
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
	case map[string]*DummyAgent:
		return any(&stage.DummyAgents_mapString).(*Type)
	case map[string]*Engine:
		return any(&stage.Engines_mapString).(*Type)
	case map[string]*Event:
		return any(&stage.Events_mapString).(*Type)
	case map[string]*GongsimCommand:
		return any(&stage.GongsimCommands_mapString).(*Type)
	case map[string]*GongsimStatus:
		return any(&stage.GongsimStatuss_mapString).(*Type)
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
	case DummyAgent:
		return any(&stage.DummyAgents).(*map[*Type]any)
	case Engine:
		return any(&stage.Engines).(*map[*Type]any)
	case Event:
		return any(&stage.Events).(*map[*Type]any)
	case GongsimCommand:
		return any(&stage.GongsimCommands).(*map[*Type]any)
	case GongsimStatus:
		return any(&stage.GongsimStatuss).(*map[*Type]any)
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
	case *DummyAgent:
		return any(&stage.DummyAgents).(*map[Type]any)
	case *Engine:
		return any(&stage.Engines).(*map[Type]any)
	case *Event:
		return any(&stage.Events).(*map[Type]any)
	case *GongsimCommand:
		return any(&stage.GongsimCommands).(*map[Type]any)
	case *GongsimStatus:
		return any(&stage.GongsimStatuss).(*map[Type]any)
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
	case DummyAgent:
		return any(&stage.DummyAgents_mapString).(*map[string]*Type)
	case Engine:
		return any(&stage.Engines_mapString).(*map[string]*Type)
	case Event:
		return any(&stage.Events_mapString).(*map[string]*Type)
	case GongsimCommand:
		return any(&stage.GongsimCommands_mapString).(*map[string]*Type)
	case GongsimStatus:
		return any(&stage.GongsimStatuss_mapString).(*map[string]*Type)
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
	case DummyAgent:
		return any(&DummyAgent{
			// Initialisation of associations
		}).(*Type)
	case Engine:
		return any(&Engine{
			// Initialisation of associations
		}).(*Type)
	case Event:
		return any(&Event{
			// Initialisation of associations
		}).(*Type)
	case GongsimCommand:
		return any(&GongsimCommand{
			// Initialisation of associations
			// field is initialized with an instance of Engine with the name of the field
			Engine: &Engine{Name: "Engine"},
		}).(*Type)
	case GongsimStatus:
		return any(&GongsimStatus{
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
	// reverse maps of direct associations of DummyAgent
	case DummyAgent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Engine
	case Engine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Event
	case Event:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongsimCommand
	case GongsimCommand:
		switch fieldname {
		// insertion point for per direct association field
		case "Engine":
			res := make(map[*Engine][]*GongsimCommand)
			for gongsimcommand := range stage.GongsimCommands {
				if gongsimcommand.Engine != nil {
					engine_ := gongsimcommand.Engine
					var gongsimcommands []*GongsimCommand
					_, ok := res[engine_]
					if ok {
						gongsimcommands = res[engine_]
					} else {
						gongsimcommands = make([]*GongsimCommand, 0)
					}
					gongsimcommands = append(gongsimcommands, gongsimcommand)
					res[engine_] = gongsimcommands
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongsimStatus
	case GongsimStatus:
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
	// reverse maps of direct associations of DummyAgent
	case DummyAgent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Engine
	case Engine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Event
	case Event:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongsimCommand
	case GongsimCommand:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongsimStatus
	case GongsimStatus:
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
	case DummyAgent:
		res = "DummyAgent"
	case Engine:
		res = "Engine"
	case Event:
		res = "Event"
	case GongsimCommand:
		res = "GongsimCommand"
	case GongsimStatus:
		res = "GongsimStatus"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *DummyAgent:
		res = "DummyAgent"
	case *Engine:
		res = "Engine"
	case *Event:
		res = "Event"
	case *GongsimCommand:
		res = "GongsimCommand"
	case *GongsimStatus:
		res = "GongsimStatus"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case DummyAgent:
		res = []string{"TechName", "Name"}
	case Engine:
		res = []string{"Name", "EndTime", "CurrentTime", "SecondsSinceStart", "Fired", "ControlMode", "State", "Speed"}
	case Event:
		res = []string{"Name", "Duration"}
	case GongsimCommand:
		res = []string{"Name", "Command", "CommandDate", "SpeedCommandType", "DateSpeedCommand", "Engine"}
	case GongsimStatus:
		res = []string{"Name", "CurrentCommand", "CompletionDate", "CurrentSpeedCommand", "SpeedCommandCompletionDate"}
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
	case DummyAgent:
		var rf ReverseField
		_ = rf
	case Engine:
		var rf ReverseField
		_ = rf
	case Event:
		var rf ReverseField
		_ = rf
	case GongsimCommand:
		var rf ReverseField
		_ = rf
	case GongsimStatus:
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
	case *DummyAgent:
		res = []string{"TechName", "Name"}
	case *Engine:
		res = []string{"Name", "EndTime", "CurrentTime", "SecondsSinceStart", "Fired", "ControlMode", "State", "Speed"}
	case *Event:
		res = []string{"Name", "Duration"}
	case *GongsimCommand:
		res = []string{"Name", "Command", "CommandDate", "SpeedCommandType", "DateSpeedCommand", "Engine"}
	case *GongsimStatus:
		res = []string{"Name", "CurrentCommand", "CompletionDate", "CurrentSpeedCommand", "SpeedCommandCompletionDate"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *DummyAgent:
		switch fieldName {
		// string value of fields
		case "TechName":
			res = inferedInstance.TechName
		case "Name":
			res = inferedInstance.Name
		}
	case *Engine:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EndTime":
			res = inferedInstance.EndTime
		case "CurrentTime":
			res = inferedInstance.CurrentTime
		case "SecondsSinceStart":
			res = fmt.Sprintf("%f", inferedInstance.SecondsSinceStart)
		case "Fired":
			res = fmt.Sprintf("%d", inferedInstance.Fired)
		case "ControlMode":
			enum := inferedInstance.ControlMode
			res = enum.ToCodeString()
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		}
	case *Event:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Duration":
			if math.Abs(inferedInstance.Duration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration.Seconds()) % 60

				if inferedInstance.Duration.Hours() < 0 {
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
				res = fmt.Sprintf("%s\n", inferedInstance.Duration.String())
			}
		}
	case *GongsimCommand:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Command":
			enum := inferedInstance.Command
			res = enum.ToCodeString()
		case "CommandDate":
			res = inferedInstance.CommandDate
		case "SpeedCommandType":
			enum := inferedInstance.SpeedCommandType
			res = enum.ToCodeString()
		case "DateSpeedCommand":
			res = inferedInstance.DateSpeedCommand
		case "Engine":
			if inferedInstance.Engine != nil {
				res = inferedInstance.Engine.Name
			}
		}
	case *GongsimStatus:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CurrentCommand":
			enum := inferedInstance.CurrentCommand
			res = enum.ToCodeString()
		case "CompletionDate":
			res = inferedInstance.CompletionDate
		case "CurrentSpeedCommand":
			enum := inferedInstance.CurrentSpeedCommand
			res = enum.ToCodeString()
		case "SpeedCommandCompletionDate":
			res = inferedInstance.SpeedCommandCompletionDate
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case DummyAgent:
		switch fieldName {
		// string value of fields
		case "TechName":
			res = inferedInstance.TechName
		case "Name":
			res = inferedInstance.Name
		}
	case Engine:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EndTime":
			res = inferedInstance.EndTime
		case "CurrentTime":
			res = inferedInstance.CurrentTime
		case "SecondsSinceStart":
			res = fmt.Sprintf("%f", inferedInstance.SecondsSinceStart)
		case "Fired":
			res = fmt.Sprintf("%d", inferedInstance.Fired)
		case "ControlMode":
			enum := inferedInstance.ControlMode
			res = enum.ToCodeString()
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		}
	case Event:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Duration":
			if math.Abs(inferedInstance.Duration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration.Seconds()) % 60

				if inferedInstance.Duration.Hours() < 0 {
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
				res = fmt.Sprintf("%s\n", inferedInstance.Duration.String())
			}
		}
	case GongsimCommand:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Command":
			enum := inferedInstance.Command
			res = enum.ToCodeString()
		case "CommandDate":
			res = inferedInstance.CommandDate
		case "SpeedCommandType":
			enum := inferedInstance.SpeedCommandType
			res = enum.ToCodeString()
		case "DateSpeedCommand":
			res = inferedInstance.DateSpeedCommand
		case "Engine":
			if inferedInstance.Engine != nil {
				res = inferedInstance.Engine.Name
			}
		}
	case GongsimStatus:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CurrentCommand":
			enum := inferedInstance.CurrentCommand
			res = enum.ToCodeString()
		case "CompletionDate":
			res = inferedInstance.CompletionDate
		case "CurrentSpeedCommand":
			enum := inferedInstance.CurrentSpeedCommand
			res = enum.ToCodeString()
		case "SpeedCommandCompletionDate":
			res = inferedInstance.SpeedCommandCompletionDate
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
