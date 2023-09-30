package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsim/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.AUTONOMOUS": ref_models.AUTONOMOUS,

	"ref_models.CHECKOUT_AGENT_STATES": ref_models.CHECKOUT_AGENT_STATES,

	"ref_models.CLIENT_CONTROL": ref_models.CLIENT_CONTROL,

	"ref_models.COMMAND_ADVANCE_10_MIN": ref_models.COMMAND_ADVANCE_10_MIN,

	"ref_models.COMMAND_DECREASE_SPEED_50_PERCENTS": ref_models.COMMAND_DECREASE_SPEED_50_PERCENTS,

	"ref_models.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE": ref_models.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE,

	"ref_models.COMMAND_FIRE_NEXT_EVENT": ref_models.COMMAND_FIRE_NEXT_EVENT,

	"ref_models.COMMAND_INCREASE_SPEED_100_PERCENTS": ref_models.COMMAND_INCREASE_SPEED_100_PERCENTS,

	"ref_models.COMMAND_PAUSE": ref_models.COMMAND_PAUSE,

	"ref_models.COMMAND_PLAY": ref_models.COMMAND_PLAY,

	"ref_models.COMMAND_RESET": ref_models.COMMAND_RESET,

	"ref_models.COMMAND_SPEED_STEADY": ref_models.COMMAND_SPEED_STEADY,

	"ref_models.COMMIT_AGENT_STATES": ref_models.COMMIT_AGENT_STATES,

	"ref_models.ControlMode": ref_models.ControlMode(""),

	"ref_models.DummyAgent": &(ref_models.DummyAgent{}),

	"ref_models.DummyAgent.Name": (ref_models.DummyAgent{}).Name,

	"ref_models.DummyAgent.TechName": (ref_models.DummyAgent{}).TechName,

	"ref_models.Engine": &(ref_models.Engine{}),

	"ref_models.Engine.ControlMode": (ref_models.Engine{}).ControlMode,

	"ref_models.Engine.CurrentTime": (ref_models.Engine{}).CurrentTime,

	"ref_models.Engine.EndTime": (ref_models.Engine{}).EndTime,

	"ref_models.Engine.Fired": (ref_models.Engine{}).Fired,

	"ref_models.Engine.Name": (ref_models.Engine{}).Name,

	"ref_models.Engine.SecondsSinceStart": (ref_models.Engine{}).SecondsSinceStart,

	"ref_models.Engine.Speed": (ref_models.Engine{}).Speed,

	"ref_models.Engine.State": (ref_models.Engine{}).State,

	"ref_models.EngineDriverState": ref_models.EngineDriverState(0),

	"ref_models.EngineRunMode": ref_models.EngineRunMode(0),

	"ref_models.EngineState": ref_models.EngineState(""),

	"ref_models.EngineStopMode": ref_models.EngineStopMode(0),

	"ref_models.Event": &(ref_models.Event{}),

	"ref_models.Event.Duration": (ref_models.Event{}).Duration,

	"ref_models.Event.Name": (ref_models.Event{}).Name,

	"ref_models.FIRE_ONE_EVENT": ref_models.FIRE_ONE_EVENT,

	"ref_models.FULL_SPEED": ref_models.FULL_SPEED,

	"ref_models.GongsimCommand": &(ref_models.GongsimCommand{}),

	"ref_models.GongsimCommand.Command": (ref_models.GongsimCommand{}).Command,

	"ref_models.GongsimCommand.CommandDate": (ref_models.GongsimCommand{}).CommandDate,

	"ref_models.GongsimCommand.DateSpeedCommand": (ref_models.GongsimCommand{}).DateSpeedCommand,

	"ref_models.GongsimCommand.Engine": (ref_models.GongsimCommand{}).Engine,

	"ref_models.GongsimCommand.Name": (ref_models.GongsimCommand{}).Name,

	"ref_models.GongsimCommand.SpeedCommandType": (ref_models.GongsimCommand{}).SpeedCommandType,

	"ref_models.GongsimCommandType": ref_models.GongsimCommandType(""),

	"ref_models.GongsimStatus": &(ref_models.GongsimStatus{}),

	"ref_models.GongsimStatus.CompletionDate": (ref_models.GongsimStatus{}).CompletionDate,

	"ref_models.GongsimStatus.CurrentCommand": (ref_models.GongsimStatus{}).CurrentCommand,

	"ref_models.GongsimStatus.CurrentSpeedCommand": (ref_models.GongsimStatus{}).CurrentSpeedCommand,

	"ref_models.GongsimStatus.Name": (ref_models.GongsimStatus{}).Name,

	"ref_models.GongsimStatus.SpeedCommandCompletionDate": (ref_models.GongsimStatus{}).SpeedCommandCompletionDate,

	"ref_models.OVER": ref_models.OVER,

	"ref_models.PAUSED": ref_models.PAUSED,

	"ref_models.RELATIVE_SPEED": ref_models.RELATIVE_SPEED,

	"ref_models.RESET_SIMULATION": ref_models.RESET_SIMULATION,

	"ref_models.RUNNING": ref_models.RUNNING,

	"ref_models.SLEEP_100_MS": ref_models.SLEEP_100_MS,

	"ref_models.STATE_CHANGED": ref_models.STATE_CHANGED,

	"ref_models.SpeedCommandType": ref_models.SpeedCommandType(""),

	"ref_models.TEN_MINUTES": ref_models.TEN_MINUTES,

	"ref_models.UNKOWN": ref_models.UNKOWN,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Button

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Command := (&models.Field{Name: `Command`}).Stage(stage)
	__Field__000001_CommandDate := (&models.Field{Name: `CommandDate`}).Stage(stage)
	__Field__000002_CompletionDate := (&models.Field{Name: `CompletionDate`}).Stage(stage)
	__Field__000003_ControlMode := (&models.Field{Name: `ControlMode`}).Stage(stage)
	__Field__000004_CurrentCommand := (&models.Field{Name: `CurrentCommand`}).Stage(stage)
	__Field__000005_CurrentSpeedCommand := (&models.Field{Name: `CurrentSpeedCommand`}).Stage(stage)
	__Field__000006_CurrentTime := (&models.Field{Name: `CurrentTime`}).Stage(stage)
	__Field__000007_DateSpeedCommand := (&models.Field{Name: `DateSpeedCommand`}).Stage(stage)
	__Field__000008_Duration := (&models.Field{Name: `Duration`}).Stage(stage)
	__Field__000009_EndTime := (&models.Field{Name: `EndTime`}).Stage(stage)
	__Field__000010_Fired := (&models.Field{Name: `Fired`}).Stage(stage)
	__Field__000011_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000014_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000015_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000016_SecondsSinceStart := (&models.Field{Name: `SecondsSinceStart`}).Stage(stage)
	__Field__000017_Speed := (&models.Field{Name: `Speed`}).Stage(stage)
	__Field__000018_SpeedCommandCompletionDate := (&models.Field{Name: `SpeedCommandCompletionDate`}).Stage(stage)
	__Field__000019_SpeedCommandType := (&models.Field{Name: `SpeedCommandType`}).Stage(stage)
	__Field__000020_State := (&models.Field{Name: `State`}).Stage(stage)
	__Field__000021_TechName := (&models.Field{Name: `TechName`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_ControlMode := (&models.GongEnumShape{Name: `NewDiagram-ControlMode`}).Stage(stage)
	__GongEnumShape__000001_NewDiagram_EngineDriverState := (&models.GongEnumShape{Name: `NewDiagram-EngineDriverState`}).Stage(stage)
	__GongEnumShape__000002_NewDiagram_EngineRunMode := (&models.GongEnumShape{Name: `NewDiagram-EngineRunMode`}).Stage(stage)
	__GongEnumShape__000003_NewDiagram_EngineState := (&models.GongEnumShape{Name: `NewDiagram-EngineState`}).Stage(stage)
	__GongEnumShape__000004_NewDiagram_EngineStopMode := (&models.GongEnumShape{Name: `NewDiagram-EngineStopMode`}).Stage(stage)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType := (&models.GongEnumShape{Name: `NewDiagram-GongsimCommandType`}).Stage(stage)
	__GongEnumShape__000006_NewDiagram_SpeedCommandType := (&models.GongEnumShape{Name: `NewDiagram-SpeedCommandType`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_AUTONOMOUS := (&models.GongEnumValueEntry{Name: `AUTONOMOUS`}).Stage(stage)
	__GongEnumValueEntry__000001_CHECKOUT_AGENT_STATES := (&models.GongEnumValueEntry{Name: `CHECKOUT_AGENT_STATES`}).Stage(stage)
	__GongEnumValueEntry__000002_CLIENT_CONTROL := (&models.GongEnumValueEntry{Name: `CLIENT_CONTROL`}).Stage(stage)
	__GongEnumValueEntry__000003_COMMAND_ADVANCE_10_MIN := (&models.GongEnumValueEntry{Name: `COMMAND_ADVANCE_10_MIN`}).Stage(stage)
	__GongEnumValueEntry__000004_COMMAND_DECREASE_SPEED_50_PERCENTS := (&models.GongEnumValueEntry{Name: `COMMAND_DECREASE_SPEED_50_PERCENTS`}).Stage(stage)
	__GongEnumValueEntry__000005_COMMAND_FIRE_EVENT_TILL_STATES_CHANGE := (&models.GongEnumValueEntry{Name: `COMMAND_FIRE_EVENT_TILL_STATES_CHANGE`}).Stage(stage)
	__GongEnumValueEntry__000006_COMMAND_FIRE_NEXT_EVENT := (&models.GongEnumValueEntry{Name: `COMMAND_FIRE_NEXT_EVENT`}).Stage(stage)
	__GongEnumValueEntry__000007_COMMAND_INCREASE_SPEED_100_PERCENTS := (&models.GongEnumValueEntry{Name: `COMMAND_INCREASE_SPEED_100_PERCENTS`}).Stage(stage)
	__GongEnumValueEntry__000008_COMMAND_PAUSE := (&models.GongEnumValueEntry{Name: `COMMAND_PAUSE`}).Stage(stage)
	__GongEnumValueEntry__000009_COMMAND_PLAY := (&models.GongEnumValueEntry{Name: `COMMAND_PLAY`}).Stage(stage)
	__GongEnumValueEntry__000010_COMMAND_RESET := (&models.GongEnumValueEntry{Name: `COMMAND_RESET`}).Stage(stage)
	__GongEnumValueEntry__000011_COMMAND_SPEED_STEADY := (&models.GongEnumValueEntry{Name: `COMMAND_SPEED_STEADY`}).Stage(stage)
	__GongEnumValueEntry__000012_COMMIT_AGENT_STATES := (&models.GongEnumValueEntry{Name: `COMMIT_AGENT_STATES`}).Stage(stage)
	__GongEnumValueEntry__000013_FIRE_ONE_EVENT := (&models.GongEnumValueEntry{Name: `FIRE_ONE_EVENT`}).Stage(stage)
	__GongEnumValueEntry__000014_FULL_SPEED := (&models.GongEnumValueEntry{Name: `FULL_SPEED`}).Stage(stage)
	__GongEnumValueEntry__000015_OVER := (&models.GongEnumValueEntry{Name: `OVER`}).Stage(stage)
	__GongEnumValueEntry__000016_PAUSED := (&models.GongEnumValueEntry{Name: `PAUSED`}).Stage(stage)
	__GongEnumValueEntry__000017_RELATIVE_SPEED := (&models.GongEnumValueEntry{Name: `RELATIVE_SPEED`}).Stage(stage)
	__GongEnumValueEntry__000018_RESET_SIMULATION := (&models.GongEnumValueEntry{Name: `RESET_SIMULATION`}).Stage(stage)
	__GongEnumValueEntry__000019_RUNNING := (&models.GongEnumValueEntry{Name: `RUNNING`}).Stage(stage)
	__GongEnumValueEntry__000020_SLEEP_100_MS := (&models.GongEnumValueEntry{Name: `SLEEP_100_MS`}).Stage(stage)
	__GongEnumValueEntry__000021_STATE_CHANGED := (&models.GongEnumValueEntry{Name: `STATE_CHANGED`}).Stage(stage)
	__GongEnumValueEntry__000022_TEN_MINUTES := (&models.GongEnumValueEntry{Name: `TEN_MINUTES`}).Stage(stage)
	__GongEnumValueEntry__000023_UNKOWN := (&models.GongEnumValueEntry{Name: `UNKOWN`}).Stage(stage)

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_DummyAgent := (&models.GongStructShape{Name: `NewDiagram-DummyAgent`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Engine := (&models.GongStructShape{Name: `NewDiagram-Engine`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_Event := (&models.GongStructShape{Name: `NewDiagram-Event`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_GongsimCommand := (&models.GongStructShape{Name: `NewDiagram-GongsimCommand`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_GongsimStatus := (&models.GongStructShape{Name: `NewDiagram-GongsimStatus`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Engine := (&models.Link{Name: `Engine`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_ControlMode := (&models.Position{Name: `Pos-NewDiagram-ControlMode`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_DummyAgent := (&models.Position{Name: `Pos-NewDiagram-DummyAgent`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Engine := (&models.Position{Name: `Pos-NewDiagram-Engine`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_EngineDriverState := (&models.Position{Name: `Pos-NewDiagram-EngineDriverState`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_EngineRunMode := (&models.Position{Name: `Pos-NewDiagram-EngineRunMode`}).Stage(stage)
	__Position__000005_Pos_NewDiagram_EngineState := (&models.Position{Name: `Pos-NewDiagram-EngineState`}).Stage(stage)
	__Position__000006_Pos_NewDiagram_EngineStopMode := (&models.Position{Name: `Pos-NewDiagram-EngineStopMode`}).Stage(stage)
	__Position__000007_Pos_NewDiagram_Event := (&models.Position{Name: `Pos-NewDiagram-Event`}).Stage(stage)
	__Position__000008_Pos_NewDiagram_GongsimCommand := (&models.Position{Name: `Pos-NewDiagram-GongsimCommand`}).Stage(stage)
	__Position__000009_Pos_NewDiagram_GongsimCommandType := (&models.Position{Name: `Pos-NewDiagram-GongsimCommandType`}).Stage(stage)
	__Position__000010_Pos_NewDiagram_GongsimStatus := (&models.Position{Name: `Pos-NewDiagram-GongsimStatus`}).Stage(stage)
	__Position__000011_Pos_NewDiagram_SpeedCommandType := (&models.Position{Name: `Pos-NewDiagram-SpeedCommandType`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-GongsimCommand and NewDiagram-Engine`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Command.Name = `Command`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.Command]
	__Field__000000_Command.Identifier = `ref_models.GongsimCommand.Command`
	__Field__000000_Command.FieldTypeAsString = ``
	__Field__000000_Command.Structname = `GongsimCommand`
	__Field__000000_Command.Fieldtypename = `GongsimCommandType`

	// Field values setup
	__Field__000001_CommandDate.Name = `CommandDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.CommandDate]
	__Field__000001_CommandDate.Identifier = `ref_models.GongsimCommand.CommandDate`
	__Field__000001_CommandDate.FieldTypeAsString = ``
	__Field__000001_CommandDate.Structname = `GongsimCommand`
	__Field__000001_CommandDate.Fieldtypename = `string`

	// Field values setup
	__Field__000002_CompletionDate.Name = `CompletionDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.CompletionDate]
	__Field__000002_CompletionDate.Identifier = `ref_models.GongsimStatus.CompletionDate`
	__Field__000002_CompletionDate.FieldTypeAsString = ``
	__Field__000002_CompletionDate.Structname = `GongsimStatus`
	__Field__000002_CompletionDate.Fieldtypename = `string`

	// Field values setup
	__Field__000003_ControlMode.Name = `ControlMode`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.ControlMode]
	__Field__000003_ControlMode.Identifier = `ref_models.Engine.ControlMode`
	__Field__000003_ControlMode.FieldTypeAsString = ``
	__Field__000003_ControlMode.Structname = `Engine`
	__Field__000003_ControlMode.Fieldtypename = `ControlMode`

	// Field values setup
	__Field__000004_CurrentCommand.Name = `CurrentCommand`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.CurrentCommand]
	__Field__000004_CurrentCommand.Identifier = `ref_models.GongsimStatus.CurrentCommand`
	__Field__000004_CurrentCommand.FieldTypeAsString = ``
	__Field__000004_CurrentCommand.Structname = `GongsimStatus`
	__Field__000004_CurrentCommand.Fieldtypename = `GongsimCommandType`

	// Field values setup
	__Field__000005_CurrentSpeedCommand.Name = `CurrentSpeedCommand`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.CurrentSpeedCommand]
	__Field__000005_CurrentSpeedCommand.Identifier = `ref_models.GongsimStatus.CurrentSpeedCommand`
	__Field__000005_CurrentSpeedCommand.FieldTypeAsString = ``
	__Field__000005_CurrentSpeedCommand.Structname = `GongsimStatus`
	__Field__000005_CurrentSpeedCommand.Fieldtypename = `SpeedCommandType`

	// Field values setup
	__Field__000006_CurrentTime.Name = `CurrentTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.CurrentTime]
	__Field__000006_CurrentTime.Identifier = `ref_models.Engine.CurrentTime`
	__Field__000006_CurrentTime.FieldTypeAsString = ``
	__Field__000006_CurrentTime.Structname = `Engine`
	__Field__000006_CurrentTime.Fieldtypename = `string`

	// Field values setup
	__Field__000007_DateSpeedCommand.Name = `DateSpeedCommand`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.DateSpeedCommand]
	__Field__000007_DateSpeedCommand.Identifier = `ref_models.GongsimCommand.DateSpeedCommand`
	__Field__000007_DateSpeedCommand.FieldTypeAsString = ``
	__Field__000007_DateSpeedCommand.Structname = `GongsimCommand`
	__Field__000007_DateSpeedCommand.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Duration.Name = `Duration`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Event.Duration]
	__Field__000008_Duration.Identifier = `ref_models.Event.Duration`
	__Field__000008_Duration.FieldTypeAsString = ``
	__Field__000008_Duration.Structname = `Event`
	__Field__000008_Duration.Fieldtypename = `Duration`

	// Field values setup
	__Field__000009_EndTime.Name = `EndTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.EndTime]
	__Field__000009_EndTime.Identifier = `ref_models.Engine.EndTime`
	__Field__000009_EndTime.FieldTypeAsString = ``
	__Field__000009_EndTime.Structname = `Engine`
	__Field__000009_EndTime.Fieldtypename = `string`

	// Field values setup
	__Field__000010_Fired.Name = `Fired`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.Fired]
	__Field__000010_Fired.Identifier = `ref_models.Engine.Fired`
	__Field__000010_Fired.FieldTypeAsString = ``
	__Field__000010_Fired.Structname = `Engine`
	__Field__000010_Fired.Fieldtypename = `int`

	// Field values setup
	__Field__000011_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.Name]
	__Field__000011_Name.Identifier = `ref_models.Engine.Name`
	__Field__000011_Name.FieldTypeAsString = ``
	__Field__000011_Name.Structname = `Engine`
	__Field__000011_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000012_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DummyAgent.Name]
	__Field__000012_Name.Identifier = `ref_models.DummyAgent.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `DummyAgent`
	__Field__000012_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000013_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Event.Name]
	__Field__000013_Name.Identifier = `ref_models.Event.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Event`
	__Field__000013_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000014_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.Name]
	__Field__000014_Name.Identifier = `ref_models.GongsimCommand.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `GongsimCommand`
	__Field__000014_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000015_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.Name]
	__Field__000015_Name.Identifier = `ref_models.GongsimStatus.Name`
	__Field__000015_Name.FieldTypeAsString = ``
	__Field__000015_Name.Structname = `GongsimStatus`
	__Field__000015_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000016_SecondsSinceStart.Name = `SecondsSinceStart`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.SecondsSinceStart]
	__Field__000016_SecondsSinceStart.Identifier = `ref_models.Engine.SecondsSinceStart`
	__Field__000016_SecondsSinceStart.FieldTypeAsString = ``
	__Field__000016_SecondsSinceStart.Structname = `Engine`
	__Field__000016_SecondsSinceStart.Fieldtypename = `float64`

	// Field values setup
	__Field__000017_Speed.Name = `Speed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.Speed]
	__Field__000017_Speed.Identifier = `ref_models.Engine.Speed`
	__Field__000017_Speed.FieldTypeAsString = ``
	__Field__000017_Speed.Structname = `Engine`
	__Field__000017_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000018_SpeedCommandCompletionDate.Name = `SpeedCommandCompletionDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.SpeedCommandCompletionDate]
	__Field__000018_SpeedCommandCompletionDate.Identifier = `ref_models.GongsimStatus.SpeedCommandCompletionDate`
	__Field__000018_SpeedCommandCompletionDate.FieldTypeAsString = ``
	__Field__000018_SpeedCommandCompletionDate.Structname = `GongsimStatus`
	__Field__000018_SpeedCommandCompletionDate.Fieldtypename = `string`

	// Field values setup
	__Field__000019_SpeedCommandType.Name = `SpeedCommandType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.SpeedCommandType]
	__Field__000019_SpeedCommandType.Identifier = `ref_models.GongsimCommand.SpeedCommandType`
	__Field__000019_SpeedCommandType.FieldTypeAsString = ``
	__Field__000019_SpeedCommandType.Structname = `GongsimCommand`
	__Field__000019_SpeedCommandType.Fieldtypename = `SpeedCommandType`

	// Field values setup
	__Field__000020_State.Name = `State`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.State]
	__Field__000020_State.Identifier = `ref_models.Engine.State`
	__Field__000020_State.FieldTypeAsString = ``
	__Field__000020_State.Structname = `Engine`
	__Field__000020_State.Fieldtypename = `EngineState`

	// Field values setup
	__Field__000021_TechName.Name = `TechName`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DummyAgent.TechName]
	__Field__000021_TechName.Identifier = `ref_models.DummyAgent.TechName`
	__Field__000021_TechName.FieldTypeAsString = ``
	__Field__000021_TechName.Structname = `DummyAgent`
	__Field__000021_TechName.Fieldtypename = `string`

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_ControlMode.Name = `NewDiagram-ControlMode`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ControlMode]
	__GongEnumShape__000000_NewDiagram_ControlMode.Identifier = `ref_models.ControlMode`
	__GongEnumShape__000000_NewDiagram_ControlMode.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_ControlMode.Heigth = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Name = `NewDiagram-EngineDriverState`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineDriverState]
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Identifier = `ref_models.EngineDriverState`
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Width = 240.000000
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Heigth = 153.000000

	// GongEnumShape values setup
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Name = `NewDiagram-EngineRunMode`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineRunMode]
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Identifier = `ref_models.EngineRunMode`
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Width = 240.000000
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Heigth = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000003_NewDiagram_EngineState.Name = `NewDiagram-EngineState`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState]
	__GongEnumShape__000003_NewDiagram_EngineState.Identifier = `ref_models.EngineState`
	__GongEnumShape__000003_NewDiagram_EngineState.Width = 240.000000
	__GongEnumShape__000003_NewDiagram_EngineState.Heigth = 108.000000

	// GongEnumShape values setup
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Name = `NewDiagram-EngineStopMode`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineStopMode]
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Identifier = `ref_models.EngineStopMode`
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Width = 451.999939
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Heigth = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Name = `NewDiagram-GongsimCommandType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommandType]
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Identifier = `ref_models.GongsimCommandType`
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Width = 450.999939
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Heigth = 153.000000

	// GongEnumShape values setup
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Name = `NewDiagram-SpeedCommandType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SpeedCommandType]
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Identifier = `ref_models.SpeedCommandType`
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Width = 435.999939
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Heigth = 108.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_AUTONOMOUS.Name = `AUTONOMOUS`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ControlMode.AUTONOMOUS]
	__GongEnumValueEntry__000000_AUTONOMOUS.Identifier = `ref_models.ControlMode.AUTONOMOUS`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000001_CHECKOUT_AGENT_STATES.Name = `CHECKOUT_AGENT_STATES`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineDriverState.CHECKOUT_AGENT_STATES]
	__GongEnumValueEntry__000001_CHECKOUT_AGENT_STATES.Identifier = `ref_models.EngineDriverState.CHECKOUT_AGENT_STATES`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000002_CLIENT_CONTROL.Name = `CLIENT_CONTROL`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ControlMode.CLIENT_CONTROL]
	__GongEnumValueEntry__000002_CLIENT_CONTROL.Identifier = `ref_models.ControlMode.CLIENT_CONTROL`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000003_COMMAND_ADVANCE_10_MIN.Name = `COMMAND_ADVANCE_10_MIN`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommandType.COMMAND_ADVANCE_10_MIN]
	__GongEnumValueEntry__000003_COMMAND_ADVANCE_10_MIN.Identifier = `ref_models.GongsimCommandType.COMMAND_ADVANCE_10_MIN`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000004_COMMAND_DECREASE_SPEED_50_PERCENTS.Name = `COMMAND_DECREASE_SPEED_50_PERCENTS`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS]
	__GongEnumValueEntry__000004_COMMAND_DECREASE_SPEED_50_PERCENTS.Identifier = `ref_models.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000005_COMMAND_FIRE_EVENT_TILL_STATES_CHANGE.Name = `COMMAND_FIRE_EVENT_TILL_STATES_CHANGE`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE]
	__GongEnumValueEntry__000005_COMMAND_FIRE_EVENT_TILL_STATES_CHANGE.Identifier = `ref_models.GongsimCommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000006_COMMAND_FIRE_NEXT_EVENT.Name = `COMMAND_FIRE_NEXT_EVENT`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommandType.COMMAND_FIRE_NEXT_EVENT]
	__GongEnumValueEntry__000006_COMMAND_FIRE_NEXT_EVENT.Identifier = `ref_models.GongsimCommandType.COMMAND_FIRE_NEXT_EVENT`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000007_COMMAND_INCREASE_SPEED_100_PERCENTS.Name = `COMMAND_INCREASE_SPEED_100_PERCENTS`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS]
	__GongEnumValueEntry__000007_COMMAND_INCREASE_SPEED_100_PERCENTS.Identifier = `ref_models.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000008_COMMAND_PAUSE.Name = `COMMAND_PAUSE`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommandType.COMMAND_PAUSE]
	__GongEnumValueEntry__000008_COMMAND_PAUSE.Identifier = `ref_models.GongsimCommandType.COMMAND_PAUSE`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000009_COMMAND_PLAY.Name = `COMMAND_PLAY`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommandType.COMMAND_PLAY]
	__GongEnumValueEntry__000009_COMMAND_PLAY.Identifier = `ref_models.GongsimCommandType.COMMAND_PLAY`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000010_COMMAND_RESET.Name = `COMMAND_RESET`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommandType.COMMAND_RESET]
	__GongEnumValueEntry__000010_COMMAND_RESET.Identifier = `ref_models.GongsimCommandType.COMMAND_RESET`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000011_COMMAND_SPEED_STEADY.Name = `COMMAND_SPEED_STEADY`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SpeedCommandType.COMMAND_SPEED_STEADY]
	__GongEnumValueEntry__000011_COMMAND_SPEED_STEADY.Identifier = `ref_models.SpeedCommandType.COMMAND_SPEED_STEADY`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000012_COMMIT_AGENT_STATES.Name = `COMMIT_AGENT_STATES`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineDriverState.COMMIT_AGENT_STATES]
	__GongEnumValueEntry__000012_COMMIT_AGENT_STATES.Identifier = `ref_models.EngineDriverState.COMMIT_AGENT_STATES`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000013_FIRE_ONE_EVENT.Name = `FIRE_ONE_EVENT`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineDriverState.FIRE_ONE_EVENT]
	__GongEnumValueEntry__000013_FIRE_ONE_EVENT.Identifier = `ref_models.EngineDriverState.FIRE_ONE_EVENT`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000014_FULL_SPEED.Name = `FULL_SPEED`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineRunMode.FULL_SPEED]
	__GongEnumValueEntry__000014_FULL_SPEED.Identifier = `ref_models.EngineRunMode.FULL_SPEED`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000015_OVER.Name = `OVER`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState.OVER]
	__GongEnumValueEntry__000015_OVER.Identifier = `ref_models.EngineState.OVER`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000016_PAUSED.Name = `PAUSED`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState.PAUSED]
	__GongEnumValueEntry__000016_PAUSED.Identifier = `ref_models.EngineState.PAUSED`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000017_RELATIVE_SPEED.Name = `RELATIVE_SPEED`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineRunMode.RELATIVE_SPEED]
	__GongEnumValueEntry__000017_RELATIVE_SPEED.Identifier = `ref_models.EngineRunMode.RELATIVE_SPEED`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000018_RESET_SIMULATION.Name = `RESET_SIMULATION`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineDriverState.RESET_SIMULATION]
	__GongEnumValueEntry__000018_RESET_SIMULATION.Identifier = `ref_models.EngineDriverState.RESET_SIMULATION`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000019_RUNNING.Name = `RUNNING`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState.RUNNING]
	__GongEnumValueEntry__000019_RUNNING.Identifier = `ref_models.EngineState.RUNNING`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000020_SLEEP_100_MS.Name = `SLEEP_100_MS`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineDriverState.SLEEP_100_MS]
	__GongEnumValueEntry__000020_SLEEP_100_MS.Identifier = `ref_models.EngineDriverState.SLEEP_100_MS`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000021_STATE_CHANGED.Name = `STATE_CHANGED`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineStopMode.STATE_CHANGED]
	__GongEnumValueEntry__000021_STATE_CHANGED.Identifier = `ref_models.EngineStopMode.STATE_CHANGED`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000022_TEN_MINUTES.Name = `TEN_MINUTES`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineStopMode.TEN_MINUTES]
	__GongEnumValueEntry__000022_TEN_MINUTES.Identifier = `ref_models.EngineStopMode.TEN_MINUTES`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000023_UNKOWN.Name = `UNKOWN`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineDriverState.UNKOWN]
	__GongEnumValueEntry__000023_UNKOWN.Identifier = `ref_models.EngineDriverState.UNKOWN`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_DummyAgent.Name = `NewDiagram-DummyAgent`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DummyAgent]
	__GongStructShape__000000_NewDiagram_DummyAgent.Identifier = `ref_models.DummyAgent`
	__GongStructShape__000000_NewDiagram_DummyAgent.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_DummyAgent.NbInstances = 0
	__GongStructShape__000000_NewDiagram_DummyAgent.Width = 240.000000
	__GongStructShape__000000_NewDiagram_DummyAgent.Heigth = 93.000000
	__GongStructShape__000000_NewDiagram_DummyAgent.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Engine.Name = `NewDiagram-Engine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine]
	__GongStructShape__000001_NewDiagram_Engine.Identifier = `ref_models.Engine`
	__GongStructShape__000001_NewDiagram_Engine.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_Engine.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Engine.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Engine.Heigth = 183.000000
	__GongStructShape__000001_NewDiagram_Engine.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Event.Name = `NewDiagram-Event`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Event]
	__GongStructShape__000002_NewDiagram_Event.Identifier = `ref_models.Event`
	__GongStructShape__000002_NewDiagram_Event.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_Event.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Event.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Event.Heigth = 93.000000
	__GongStructShape__000002_NewDiagram_Event.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_GongsimCommand.Name = `NewDiagram-GongsimCommand`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand]
	__GongStructShape__000003_NewDiagram_GongsimCommand.Identifier = `ref_models.GongsimCommand`
	__GongStructShape__000003_NewDiagram_GongsimCommand.ShowNbInstances = true
	__GongStructShape__000003_NewDiagram_GongsimCommand.NbInstances = 0
	__GongStructShape__000003_NewDiagram_GongsimCommand.Width = 389.000000
	__GongStructShape__000003_NewDiagram_GongsimCommand.Heigth = 138.000000
	__GongStructShape__000003_NewDiagram_GongsimCommand.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_GongsimStatus.Name = `NewDiagram-GongsimStatus`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus]
	__GongStructShape__000004_NewDiagram_GongsimStatus.Identifier = `ref_models.GongsimStatus`
	__GongStructShape__000004_NewDiagram_GongsimStatus.ShowNbInstances = true
	__GongStructShape__000004_NewDiagram_GongsimStatus.NbInstances = 0
	__GongStructShape__000004_NewDiagram_GongsimStatus.Width = 365.000000
	__GongStructShape__000004_NewDiagram_GongsimStatus.Heigth = 138.000000
	__GongStructShape__000004_NewDiagram_GongsimStatus.IsSelected = false

	// Link values setup
	__Link__000000_Engine.Name = `Engine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.Engine]
	__Link__000000_Engine.Identifier = `ref_models.GongsimCommand.Engine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine]
	__Link__000000_Engine.Fieldtypename = `ref_models.Engine`
	__Link__000000_Engine.FieldOffsetX = 15.000000
	__Link__000000_Engine.FieldOffsetY = -19.000000
	__Link__000000_Engine.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Engine.TargetMultiplicityOffsetX = 13.000000
	__Link__000000_Engine.TargetMultiplicityOffsetY = 29.000000
	__Link__000000_Engine.SourceMultiplicity = models.MANY
	__Link__000000_Engine.SourceMultiplicityOffsetX = 9.000000
	__Link__000000_Engine.SourceMultiplicityOffsetY = 34.000000
	__Link__000000_Engine.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Engine.StartRatio = 0.500000
	__Link__000000_Engine.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Engine.EndRatio = 0.500000
	__Link__000000_Engine.CornerOffsetRatio = 1.052099

	// Position values setup
	__Position__000000_Pos_NewDiagram_ControlMode.X = 798.000000
	__Position__000000_Pos_NewDiagram_ControlMode.Y = 783.000000
	__Position__000000_Pos_NewDiagram_ControlMode.Name = `Pos-NewDiagram-ControlMode`

	// Position values setup
	__Position__000001_Pos_NewDiagram_DummyAgent.X = 50.000000
	__Position__000001_Pos_NewDiagram_DummyAgent.Y = 670.000000
	__Position__000001_Pos_NewDiagram_DummyAgent.Name = `Pos-NewDiagram-DummyAgent`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Engine.X = 44.000000
	__Position__000002_Pos_NewDiagram_Engine.Y = 85.000000
	__Position__000002_Pos_NewDiagram_Engine.Name = `Pos-NewDiagram-Engine`

	// Position values setup
	__Position__000003_Pos_NewDiagram_EngineDriverState.X = 785.000000
	__Position__000003_Pos_NewDiagram_EngineDriverState.Y = 589.000000
	__Position__000003_Pos_NewDiagram_EngineDriverState.Name = `Pos-NewDiagram-EngineDriverState`

	// Position values setup
	__Position__000004_Pos_NewDiagram_EngineRunMode.X = 476.000000
	__Position__000004_Pos_NewDiagram_EngineRunMode.Y = 259.000000
	__Position__000004_Pos_NewDiagram_EngineRunMode.Name = `Pos-NewDiagram-EngineRunMode`

	// Position values setup
	__Position__000005_Pos_NewDiagram_EngineState.X = 488.000000
	__Position__000005_Pos_NewDiagram_EngineState.Y = 86.000000
	__Position__000005_Pos_NewDiagram_EngineState.Name = `Pos-NewDiagram-EngineState`

	// Position values setup
	__Position__000006_Pos_NewDiagram_EngineStopMode.X = 780.000000
	__Position__000006_Pos_NewDiagram_EngineStopMode.Y = 431.000000
	__Position__000006_Pos_NewDiagram_EngineStopMode.Name = `Pos-NewDiagram-EngineStopMode`

	// Position values setup
	__Position__000007_Pos_NewDiagram_Event.X = 60.000000
	__Position__000007_Pos_NewDiagram_Event.Y = 520.000000
	__Position__000007_Pos_NewDiagram_Event.Name = `Pos-NewDiagram-Event`

	// Position values setup
	__Position__000008_Pos_NewDiagram_GongsimCommand.X = 29.000000
	__Position__000008_Pos_NewDiagram_GongsimCommand.Y = 318.000000
	__Position__000008_Pos_NewDiagram_GongsimCommand.Name = `Pos-NewDiagram-GongsimCommand`

	// Position values setup
	__Position__000009_Pos_NewDiagram_GongsimCommandType.X = 778.000000
	__Position__000009_Pos_NewDiagram_GongsimCommandType.Y = 236.000000
	__Position__000009_Pos_NewDiagram_GongsimCommandType.Name = `Pos-NewDiagram-GongsimCommandType`

	// Position values setup
	__Position__000010_Pos_NewDiagram_GongsimStatus.X = 367.000000
	__Position__000010_Pos_NewDiagram_GongsimStatus.Y = 534.000000
	__Position__000010_Pos_NewDiagram_GongsimStatus.Name = `Pos-NewDiagram-GongsimStatus`

	// Position values setup
	__Position__000011_Pos_NewDiagram_SpeedCommandType.X = 777.000000
	__Position__000011_Pos_NewDiagram_SpeedCommandType.Y = 83.000000
	__Position__000011_Pos_NewDiagram_SpeedCommandType.Name = `Pos-NewDiagram-SpeedCommandType`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine.X = 644.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine.Y = 274.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-GongsimCommand and NewDiagram-Engine`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Engine)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_GongsimCommand)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_GongsimStatus)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_DummyAgent)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Event)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000003_NewDiagram_EngineState)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000002_NewDiagram_EngineRunMode)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000001_NewDiagram_EngineDriverState)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000006_NewDiagram_SpeedCommandType)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000005_NewDiagram_GongsimCommandType)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000004_NewDiagram_EngineStopMode)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000000_NewDiagram_ControlMode)
	__GongEnumShape__000000_NewDiagram_ControlMode.Position = __Position__000000_Pos_NewDiagram_ControlMode
	__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys, __GongEnumValueEntry__000000_AUTONOMOUS)
	__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys, __GongEnumValueEntry__000002_CLIENT_CONTROL)
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Position = __Position__000003_Pos_NewDiagram_EngineDriverState
	__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys, __GongEnumValueEntry__000012_COMMIT_AGENT_STATES)
	__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys, __GongEnumValueEntry__000001_CHECKOUT_AGENT_STATES)
	__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys, __GongEnumValueEntry__000013_FIRE_ONE_EVENT)
	__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys, __GongEnumValueEntry__000020_SLEEP_100_MS)
	__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys, __GongEnumValueEntry__000018_RESET_SIMULATION)
	__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_EngineDriverState.GongEnumValueEntrys, __GongEnumValueEntry__000023_UNKOWN)
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Position = __Position__000004_Pos_NewDiagram_EngineRunMode
	__GongEnumShape__000002_NewDiagram_EngineRunMode.GongEnumValueEntrys = append(__GongEnumShape__000002_NewDiagram_EngineRunMode.GongEnumValueEntrys, __GongEnumValueEntry__000017_RELATIVE_SPEED)
	__GongEnumShape__000002_NewDiagram_EngineRunMode.GongEnumValueEntrys = append(__GongEnumShape__000002_NewDiagram_EngineRunMode.GongEnumValueEntrys, __GongEnumValueEntry__000014_FULL_SPEED)
	__GongEnumShape__000003_NewDiagram_EngineState.Position = __Position__000005_Pos_NewDiagram_EngineState
	__GongEnumShape__000003_NewDiagram_EngineState.GongEnumValueEntrys = append(__GongEnumShape__000003_NewDiagram_EngineState.GongEnumValueEntrys, __GongEnumValueEntry__000019_RUNNING)
	__GongEnumShape__000003_NewDiagram_EngineState.GongEnumValueEntrys = append(__GongEnumShape__000003_NewDiagram_EngineState.GongEnumValueEntrys, __GongEnumValueEntry__000016_PAUSED)
	__GongEnumShape__000003_NewDiagram_EngineState.GongEnumValueEntrys = append(__GongEnumShape__000003_NewDiagram_EngineState.GongEnumValueEntrys, __GongEnumValueEntry__000015_OVER)
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Position = __Position__000006_Pos_NewDiagram_EngineStopMode
	__GongEnumShape__000004_NewDiagram_EngineStopMode.GongEnumValueEntrys = append(__GongEnumShape__000004_NewDiagram_EngineStopMode.GongEnumValueEntrys, __GongEnumValueEntry__000022_TEN_MINUTES)
	__GongEnumShape__000004_NewDiagram_EngineStopMode.GongEnumValueEntrys = append(__GongEnumShape__000004_NewDiagram_EngineStopMode.GongEnumValueEntrys, __GongEnumValueEntry__000021_STATE_CHANGED)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Position = __Position__000009_Pos_NewDiagram_GongsimCommandType
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys = append(__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000009_COMMAND_PLAY)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys = append(__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000008_COMMAND_PAUSE)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys = append(__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000006_COMMAND_FIRE_NEXT_EVENT)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys = append(__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000005_COMMAND_FIRE_EVENT_TILL_STATES_CHANGE)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys = append(__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000010_COMMAND_RESET)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys = append(__GongEnumShape__000005_NewDiagram_GongsimCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000003_COMMAND_ADVANCE_10_MIN)
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Position = __Position__000011_Pos_NewDiagram_SpeedCommandType
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.GongEnumValueEntrys = append(__GongEnumShape__000006_NewDiagram_SpeedCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000007_COMMAND_INCREASE_SPEED_100_PERCENTS)
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.GongEnumValueEntrys = append(__GongEnumShape__000006_NewDiagram_SpeedCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000004_COMMAND_DECREASE_SPEED_50_PERCENTS)
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.GongEnumValueEntrys = append(__GongEnumShape__000006_NewDiagram_SpeedCommandType.GongEnumValueEntrys, __GongEnumValueEntry__000011_COMMAND_SPEED_STEADY)
	__GongStructShape__000000_NewDiagram_DummyAgent.Position = __Position__000001_Pos_NewDiagram_DummyAgent
	__GongStructShape__000000_NewDiagram_DummyAgent.Fields = append(__GongStructShape__000000_NewDiagram_DummyAgent.Fields, __Field__000021_TechName)
	__GongStructShape__000000_NewDiagram_DummyAgent.Fields = append(__GongStructShape__000000_NewDiagram_DummyAgent.Fields, __Field__000012_Name)
	__GongStructShape__000001_NewDiagram_Engine.Position = __Position__000002_Pos_NewDiagram_Engine
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000011_Name)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000009_EndTime)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000006_CurrentTime)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000016_SecondsSinceStart)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000010_Fired)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000003_ControlMode)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000020_State)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000017_Speed)
	__GongStructShape__000002_NewDiagram_Event.Position = __Position__000007_Pos_NewDiagram_Event
	__GongStructShape__000002_NewDiagram_Event.Fields = append(__GongStructShape__000002_NewDiagram_Event.Fields, __Field__000013_Name)
	__GongStructShape__000002_NewDiagram_Event.Fields = append(__GongStructShape__000002_NewDiagram_Event.Fields, __Field__000008_Duration)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Position = __Position__000008_Pos_NewDiagram_GongsimCommand
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000014_Name)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000000_Command)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000001_CommandDate)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000019_SpeedCommandType)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000007_DateSpeedCommand)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Links = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Links, __Link__000000_Engine)
	__GongStructShape__000004_NewDiagram_GongsimStatus.Position = __Position__000010_Pos_NewDiagram_GongsimStatus
	__GongStructShape__000004_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000004_NewDiagram_GongsimStatus.Fields, __Field__000015_Name)
	__GongStructShape__000004_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000004_NewDiagram_GongsimStatus.Fields, __Field__000004_CurrentCommand)
	__GongStructShape__000004_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000004_NewDiagram_GongsimStatus.Fields, __Field__000002_CompletionDate)
	__GongStructShape__000004_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000004_NewDiagram_GongsimStatus.Fields, __Field__000005_CurrentSpeedCommand)
	__GongStructShape__000004_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000004_NewDiagram_GongsimStatus.Fields, __Field__000018_SpeedCommandCompletionDate)
	__Link__000000_Engine.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine
}


