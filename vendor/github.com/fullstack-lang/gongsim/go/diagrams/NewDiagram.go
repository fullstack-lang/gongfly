package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsim/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_Command := (&models.Field{}).Stage(stage)
	__Field__000001_CommandDate := (&models.Field{}).Stage(stage)
	__Field__000002_CompletionDate := (&models.Field{}).Stage(stage)
	__Field__000003_ControlMode := (&models.Field{}).Stage(stage)
	__Field__000004_CurrentCommand := (&models.Field{}).Stage(stage)
	__Field__000005_CurrentSpeedCommand := (&models.Field{}).Stage(stage)
	__Field__000006_CurrentTime := (&models.Field{}).Stage(stage)
	__Field__000007_DateSpeedCommand := (&models.Field{}).Stage(stage)
	__Field__000008_Duration := (&models.Field{}).Stage(stage)
	__Field__000009_EndTime := (&models.Field{}).Stage(stage)
	__Field__000010_Fired := (&models.Field{}).Stage(stage)
	__Field__000011_Name := (&models.Field{}).Stage(stage)
	__Field__000012_Name := (&models.Field{}).Stage(stage)
	__Field__000013_Name := (&models.Field{}).Stage(stage)
	__Field__000014_Name := (&models.Field{}).Stage(stage)
	__Field__000015_Name := (&models.Field{}).Stage(stage)
	__Field__000016_SecondsSinceStart := (&models.Field{}).Stage(stage)
	__Field__000017_Speed := (&models.Field{}).Stage(stage)
	__Field__000018_SpeedCommandCompletionDate := (&models.Field{}).Stage(stage)
	__Field__000019_SpeedCommandType := (&models.Field{}).Stage(stage)
	__Field__000020_State := (&models.Field{}).Stage(stage)
	__Field__000021_TechName := (&models.Field{}).Stage(stage)

	__GongEnumShape__000000_NewDiagram_ControlMode := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000001_NewDiagram_EngineDriverState := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000002_NewDiagram_EngineRunMode := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000003_NewDiagram_EngineState := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000004_NewDiagram_EngineStopMode := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000005_NewDiagram_GongsimCommandType := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000006_NewDiagram_SpeedCommandType := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueEntry__000000_AUTONOMOUS := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000001_CHECKOUT_AGENT_STATES := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000002_CLIENT_CONTROL := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000003_COMMAND_ADVANCE_10_MIN := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000004_COMMAND_DECREASE_SPEED_50_PERCENTS := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000005_COMMAND_FIRE_EVENT_TILL_STATES_CHANGE := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000006_COMMAND_FIRE_NEXT_EVENT := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000007_COMMAND_INCREASE_SPEED_100_PERCENTS := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000008_COMMAND_PAUSE := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000009_COMMAND_PLAY := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000010_COMMAND_RESET := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000011_COMMAND_SPEED_STEADY := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000012_COMMIT_AGENT_STATES := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000013_FIRE_ONE_EVENT := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000014_FULL_SPEED := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000015_OVER := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000016_PAUSED := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000017_RELATIVE_SPEED := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000018_RESET_SIMULATION := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000019_RUNNING := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000020_SLEEP_100_MS := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000021_STATE_CHANGED := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000022_TEN_MINUTES := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000023_UNKOWN := (&models.GongEnumValueEntry{}).Stage(stage)

	__GongStructShape__000000_NewDiagram_DummyAgent := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Engine := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_NewDiagram_Event := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_NewDiagram_GongsimCommand := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_NewDiagram_Status := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Engine := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_NewDiagram_ControlMode := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_NewDiagram_DummyAgent := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Engine := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_NewDiagram_EngineDriverState := (&models.Position{}).Stage(stage)
	__Position__000004_Pos_NewDiagram_EngineRunMode := (&models.Position{}).Stage(stage)
	__Position__000005_Pos_NewDiagram_EngineState := (&models.Position{}).Stage(stage)
	__Position__000006_Pos_NewDiagram_EngineStopMode := (&models.Position{}).Stage(stage)
	__Position__000007_Pos_NewDiagram_Event := (&models.Position{}).Stage(stage)
	__Position__000008_Pos_NewDiagram_GongsimCommand := (&models.Position{}).Stage(stage)
	__Position__000009_Pos_NewDiagram_GongsimCommandType := (&models.Position{}).Stage(stage)
	__Position__000010_Pos_NewDiagram_Status := (&models.Position{}).Stage(stage)
	__Position__000011_Pos_NewDiagram_SpeedCommandType := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = false

	__Field__000000_Command.Name = `Command`

	//gong:ident [ref_models.Command.Command] comment added to overcome the problem with the comment map association
	__Field__000000_Command.Identifier = `ref_models.Command.Command`
	__Field__000000_Command.FieldTypeAsString = ``
	__Field__000000_Command.Structname = `Command`
	__Field__000000_Command.Fieldtypename = `GongsimCommandType`

	__Field__000001_CommandDate.Name = `CommandDate`

	//gong:ident [ref_models.Command.CommandDate] comment added to overcome the problem with the comment map association
	__Field__000001_CommandDate.Identifier = `ref_models.Command.CommandDate`
	__Field__000001_CommandDate.FieldTypeAsString = ``
	__Field__000001_CommandDate.Structname = `Command`
	__Field__000001_CommandDate.Fieldtypename = `string`

	__Field__000002_CompletionDate.Name = `CompletionDate`

	//gong:ident [ref_models.Status.CompletionDate] comment added to overcome the problem with the comment map association
	__Field__000002_CompletionDate.Identifier = `ref_models.Status.CompletionDate`
	__Field__000002_CompletionDate.FieldTypeAsString = ``
	__Field__000002_CompletionDate.Structname = `Status`
	__Field__000002_CompletionDate.Fieldtypename = `string`

	__Field__000003_ControlMode.Name = `ControlMode`

	//gong:ident [ref_models.Engine.ControlMode] comment added to overcome the problem with the comment map association
	__Field__000003_ControlMode.Identifier = `ref_models.Engine.ControlMode`
	__Field__000003_ControlMode.FieldTypeAsString = ``
	__Field__000003_ControlMode.Structname = `Engine`
	__Field__000003_ControlMode.Fieldtypename = `ControlMode`

	__Field__000004_CurrentCommand.Name = `CurrentCommand`

	//gong:ident [ref_models.Status.CurrentCommand] comment added to overcome the problem with the comment map association
	__Field__000004_CurrentCommand.Identifier = `ref_models.Status.CurrentCommand`
	__Field__000004_CurrentCommand.FieldTypeAsString = ``
	__Field__000004_CurrentCommand.Structname = `Status`
	__Field__000004_CurrentCommand.Fieldtypename = `GongsimCommandType`

	__Field__000005_CurrentSpeedCommand.Name = `CurrentSpeedCommand`

	//gong:ident [ref_models.Status.CurrentSpeedCommand] comment added to overcome the problem with the comment map association
	__Field__000005_CurrentSpeedCommand.Identifier = `ref_models.Status.CurrentSpeedCommand`
	__Field__000005_CurrentSpeedCommand.FieldTypeAsString = ``
	__Field__000005_CurrentSpeedCommand.Structname = `Status`
	__Field__000005_CurrentSpeedCommand.Fieldtypename = `SpeedCommandType`

	__Field__000006_CurrentTime.Name = `CurrentTime`

	//gong:ident [ref_models.Engine.CurrentTime] comment added to overcome the problem with the comment map association
	__Field__000006_CurrentTime.Identifier = `ref_models.Engine.CurrentTime`
	__Field__000006_CurrentTime.FieldTypeAsString = ``
	__Field__000006_CurrentTime.Structname = `Engine`
	__Field__000006_CurrentTime.Fieldtypename = `string`

	__Field__000007_DateSpeedCommand.Name = `DateSpeedCommand`

	//gong:ident [ref_models.Command.DateSpeedCommand] comment added to overcome the problem with the comment map association
	__Field__000007_DateSpeedCommand.Identifier = `ref_models.Command.DateSpeedCommand`
	__Field__000007_DateSpeedCommand.FieldTypeAsString = ``
	__Field__000007_DateSpeedCommand.Structname = `Command`
	__Field__000007_DateSpeedCommand.Fieldtypename = `string`

	__Field__000008_Duration.Name = `Duration`

	//gong:ident [ref_models.Event.Duration] comment added to overcome the problem with the comment map association
	__Field__000008_Duration.Identifier = `ref_models.Event.Duration`
	__Field__000008_Duration.FieldTypeAsString = ``
	__Field__000008_Duration.Structname = `Event`
	__Field__000008_Duration.Fieldtypename = `Duration`

	__Field__000009_EndTime.Name = `EndTime`

	//gong:ident [ref_models.Engine.EndTime] comment added to overcome the problem with the comment map association
	__Field__000009_EndTime.Identifier = `ref_models.Engine.EndTime`
	__Field__000009_EndTime.FieldTypeAsString = ``
	__Field__000009_EndTime.Structname = `Engine`
	__Field__000009_EndTime.Fieldtypename = `string`

	__Field__000010_Fired.Name = `Fired`

	//gong:ident [ref_models.Engine.Fired] comment added to overcome the problem with the comment map association
	__Field__000010_Fired.Identifier = `ref_models.Engine.Fired`
	__Field__000010_Fired.FieldTypeAsString = ``
	__Field__000010_Fired.Structname = `Engine`
	__Field__000010_Fired.Fieldtypename = `int`

	__Field__000011_Name.Name = `Name`

	//gong:ident [ref_models.DummyAgent.Name] comment added to overcome the problem with the comment map association
	__Field__000011_Name.Identifier = `ref_models.DummyAgent.Name`
	__Field__000011_Name.FieldTypeAsString = ``
	__Field__000011_Name.Structname = `DummyAgent`
	__Field__000011_Name.Fieldtypename = `string`

	__Field__000012_Name.Name = `Name`

	//gong:ident [ref_models.Engine.Name] comment added to overcome the problem with the comment map association
	__Field__000012_Name.Identifier = `ref_models.Engine.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `Engine`
	__Field__000012_Name.Fieldtypename = `string`

	__Field__000013_Name.Name = `Name`

	//gong:ident [ref_models.Status.Name] comment added to overcome the problem with the comment map association
	__Field__000013_Name.Identifier = `ref_models.Status.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Status`
	__Field__000013_Name.Fieldtypename = `string`

	__Field__000014_Name.Name = `Name`

	//gong:ident [ref_models.Event.Name] comment added to overcome the problem with the comment map association
	__Field__000014_Name.Identifier = `ref_models.Event.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `Event`
	__Field__000014_Name.Fieldtypename = `string`

	__Field__000015_Name.Name = `Name`

	//gong:ident [ref_models.Command.Name] comment added to overcome the problem with the comment map association
	__Field__000015_Name.Identifier = `ref_models.Command.Name`
	__Field__000015_Name.FieldTypeAsString = ``
	__Field__000015_Name.Structname = `Command`
	__Field__000015_Name.Fieldtypename = `string`

	__Field__000016_SecondsSinceStart.Name = `SecondsSinceStart`

	//gong:ident [ref_models.Engine.SecondsSinceStart] comment added to overcome the problem with the comment map association
	__Field__000016_SecondsSinceStart.Identifier = `ref_models.Engine.SecondsSinceStart`
	__Field__000016_SecondsSinceStart.FieldTypeAsString = ``
	__Field__000016_SecondsSinceStart.Structname = `Engine`
	__Field__000016_SecondsSinceStart.Fieldtypename = `float64`

	__Field__000017_Speed.Name = `Speed`

	//gong:ident [ref_models.Engine.Speed] comment added to overcome the problem with the comment map association
	__Field__000017_Speed.Identifier = `ref_models.Engine.Speed`
	__Field__000017_Speed.FieldTypeAsString = ``
	__Field__000017_Speed.Structname = `Engine`
	__Field__000017_Speed.Fieldtypename = `float64`

	__Field__000018_SpeedCommandCompletionDate.Name = `SpeedCommandCompletionDate`

	//gong:ident [ref_models.Status.SpeedCommandCompletionDate] comment added to overcome the problem with the comment map association
	__Field__000018_SpeedCommandCompletionDate.Identifier = `ref_models.Status.SpeedCommandCompletionDate`
	__Field__000018_SpeedCommandCompletionDate.FieldTypeAsString = ``
	__Field__000018_SpeedCommandCompletionDate.Structname = `Status`
	__Field__000018_SpeedCommandCompletionDate.Fieldtypename = `string`

	__Field__000019_SpeedCommandType.Name = `SpeedCommandType`

	//gong:ident [ref_models.Command.SpeedCommandType] comment added to overcome the problem with the comment map association
	__Field__000019_SpeedCommandType.Identifier = `ref_models.Command.SpeedCommandType`
	__Field__000019_SpeedCommandType.FieldTypeAsString = ``
	__Field__000019_SpeedCommandType.Structname = `Command`
	__Field__000019_SpeedCommandType.Fieldtypename = `SpeedCommandType`

	__Field__000020_State.Name = `State`

	//gong:ident [ref_models.Engine.State] comment added to overcome the problem with the comment map association
	__Field__000020_State.Identifier = `ref_models.Engine.State`
	__Field__000020_State.FieldTypeAsString = ``
	__Field__000020_State.Structname = `Engine`
	__Field__000020_State.Fieldtypename = `EngineState`

	__Field__000021_TechName.Name = `TechName`

	//gong:ident [ref_models.DummyAgent.TechName] comment added to overcome the problem with the comment map association
	__Field__000021_TechName.Identifier = `ref_models.DummyAgent.TechName`
	__Field__000021_TechName.FieldTypeAsString = ``
	__Field__000021_TechName.Structname = `DummyAgent`
	__Field__000021_TechName.Fieldtypename = `string`

	__GongEnumShape__000000_NewDiagram_ControlMode.Name = `NewDiagram-ControlMode`

	//gong:ident [ref_models.ControlMode] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_NewDiagram_ControlMode.Identifier = `ref_models.ControlMode`
	__GongEnumShape__000000_NewDiagram_ControlMode.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_ControlMode.Height = 93.000000

	__GongEnumShape__000001_NewDiagram_EngineDriverState.Name = `NewDiagram-EngineDriverState`

	//gong:ident [ref_models.EngineDriverState] comment added to overcome the problem with the comment map association
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Identifier = `ref_models.EngineDriverState`
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Width = 240.000000
	__GongEnumShape__000001_NewDiagram_EngineDriverState.Height = 153.000000

	__GongEnumShape__000002_NewDiagram_EngineRunMode.Name = `NewDiagram-EngineRunMode`

	//gong:ident [ref_models.EngineRunMode] comment added to overcome the problem with the comment map association
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Identifier = `ref_models.EngineRunMode`
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Width = 240.000000
	__GongEnumShape__000002_NewDiagram_EngineRunMode.Height = 93.000000

	__GongEnumShape__000003_NewDiagram_EngineState.Name = `NewDiagram-EngineState`

	//gong:ident [ref_models.EngineState] comment added to overcome the problem with the comment map association
	__GongEnumShape__000003_NewDiagram_EngineState.Identifier = `ref_models.EngineState`
	__GongEnumShape__000003_NewDiagram_EngineState.Width = 240.000000
	__GongEnumShape__000003_NewDiagram_EngineState.Height = 108.000000

	__GongEnumShape__000004_NewDiagram_EngineStopMode.Name = `NewDiagram-EngineStopMode`

	//gong:ident [ref_models.EngineStopMode] comment added to overcome the problem with the comment map association
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Identifier = `ref_models.EngineStopMode`
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Width = 451.999939
	__GongEnumShape__000004_NewDiagram_EngineStopMode.Height = 93.000000

	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Name = `NewDiagram-GongsimCommandType`

	//gong:ident [ref_models.CommandType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Identifier = `ref_models.CommandType`
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Width = 450.999939
	__GongEnumShape__000005_NewDiagram_GongsimCommandType.Height = 153.000000

	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Name = `NewDiagram-SpeedCommandType`

	//gong:ident [ref_models.SpeedCommandType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Identifier = `ref_models.SpeedCommandType`
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Width = 435.999939
	__GongEnumShape__000006_NewDiagram_SpeedCommandType.Height = 108.000000

	__GongEnumValueEntry__000000_AUTONOMOUS.Name = `AUTONOMOUS`

	//gong:ident [ref_models.ControlMode.AUTONOMOUS] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000000_AUTONOMOUS.Identifier = `ref_models.ControlMode.AUTONOMOUS`

	__GongEnumValueEntry__000001_CHECKOUT_AGENT_STATES.Name = `CHECKOUT_AGENT_STATES`

	//gong:ident [ref_models.EngineDriverState.CHECKOUT_AGENT_STATES] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000001_CHECKOUT_AGENT_STATES.Identifier = `ref_models.EngineDriverState.CHECKOUT_AGENT_STATES`

	__GongEnumValueEntry__000002_CLIENT_CONTROL.Name = `CLIENT_CONTROL`

	//gong:ident [ref_models.ControlMode.CLIENT_CONTROL] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000002_CLIENT_CONTROL.Identifier = `ref_models.ControlMode.CLIENT_CONTROL`

	__GongEnumValueEntry__000003_COMMAND_ADVANCE_10_MIN.Name = `COMMAND_ADVANCE_10_MIN`

	//gong:ident [ref_models.CommandType.COMMAND_ADVANCE_10_MIN] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000003_COMMAND_ADVANCE_10_MIN.Identifier = `ref_models.CommandType.COMMAND_ADVANCE_10_MIN`

	__GongEnumValueEntry__000004_COMMAND_DECREASE_SPEED_50_PERCENTS.Name = `COMMAND_DECREASE_SPEED_50_PERCENTS`

	//gong:ident [ref_models.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000004_COMMAND_DECREASE_SPEED_50_PERCENTS.Identifier = `ref_models.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS`

	__GongEnumValueEntry__000005_COMMAND_FIRE_EVENT_TILL_STATES_CHANGE.Name = `COMMAND_FIRE_EVENT_TILL_STATES_CHANGE`

	//gong:ident [ref_models.CommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000005_COMMAND_FIRE_EVENT_TILL_STATES_CHANGE.Identifier = `ref_models.CommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE`

	__GongEnumValueEntry__000006_COMMAND_FIRE_NEXT_EVENT.Name = `COMMAND_FIRE_NEXT_EVENT`

	//gong:ident [ref_models.CommandType.COMMAND_FIRE_NEXT_EVENT] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000006_COMMAND_FIRE_NEXT_EVENT.Identifier = `ref_models.CommandType.COMMAND_FIRE_NEXT_EVENT`

	__GongEnumValueEntry__000007_COMMAND_INCREASE_SPEED_100_PERCENTS.Name = `COMMAND_INCREASE_SPEED_100_PERCENTS`

	//gong:ident [ref_models.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000007_COMMAND_INCREASE_SPEED_100_PERCENTS.Identifier = `ref_models.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS`

	__GongEnumValueEntry__000008_COMMAND_PAUSE.Name = `COMMAND_PAUSE`

	//gong:ident [ref_models.CommandType.COMMAND_PAUSE] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000008_COMMAND_PAUSE.Identifier = `ref_models.CommandType.COMMAND_PAUSE`

	__GongEnumValueEntry__000009_COMMAND_PLAY.Name = `COMMAND_PLAY`

	//gong:ident [ref_models.CommandType.COMMAND_PLAY] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000009_COMMAND_PLAY.Identifier = `ref_models.CommandType.COMMAND_PLAY`

	__GongEnumValueEntry__000010_COMMAND_RESET.Name = `COMMAND_RESET`

	//gong:ident [ref_models.CommandType.COMMAND_RESET] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000010_COMMAND_RESET.Identifier = `ref_models.CommandType.COMMAND_RESET`

	__GongEnumValueEntry__000011_COMMAND_SPEED_STEADY.Name = `COMMAND_SPEED_STEADY`

	//gong:ident [ref_models.SpeedCommandType.COMMAND_SPEED_STEADY] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000011_COMMAND_SPEED_STEADY.Identifier = `ref_models.SpeedCommandType.COMMAND_SPEED_STEADY`

	__GongEnumValueEntry__000012_COMMIT_AGENT_STATES.Name = `COMMIT_AGENT_STATES`

	//gong:ident [ref_models.EngineDriverState.COMMIT_AGENT_STATES] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000012_COMMIT_AGENT_STATES.Identifier = `ref_models.EngineDriverState.COMMIT_AGENT_STATES`

	__GongEnumValueEntry__000013_FIRE_ONE_EVENT.Name = `FIRE_ONE_EVENT`

	//gong:ident [ref_models.EngineDriverState.FIRE_ONE_EVENT] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000013_FIRE_ONE_EVENT.Identifier = `ref_models.EngineDriverState.FIRE_ONE_EVENT`

	__GongEnumValueEntry__000014_FULL_SPEED.Name = `FULL_SPEED`

	//gong:ident [ref_models.EngineRunMode.FULL_SPEED] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000014_FULL_SPEED.Identifier = `ref_models.EngineRunMode.FULL_SPEED`

	__GongEnumValueEntry__000015_OVER.Name = `OVER`

	//gong:ident [ref_models.EngineState.OVER] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000015_OVER.Identifier = `ref_models.EngineState.OVER`

	__GongEnumValueEntry__000016_PAUSED.Name = `PAUSED`

	//gong:ident [ref_models.EngineState.PAUSED] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000016_PAUSED.Identifier = `ref_models.EngineState.PAUSED`

	__GongEnumValueEntry__000017_RELATIVE_SPEED.Name = `RELATIVE_SPEED`

	//gong:ident [ref_models.EngineRunMode.RELATIVE_SPEED] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000017_RELATIVE_SPEED.Identifier = `ref_models.EngineRunMode.RELATIVE_SPEED`

	__GongEnumValueEntry__000018_RESET_SIMULATION.Name = `RESET_SIMULATION`

	//gong:ident [ref_models.EngineDriverState.RESET_SIMULATION] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000018_RESET_SIMULATION.Identifier = `ref_models.EngineDriverState.RESET_SIMULATION`

	__GongEnumValueEntry__000019_RUNNING.Name = `RUNNING`

	//gong:ident [ref_models.EngineState.RUNNING] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000019_RUNNING.Identifier = `ref_models.EngineState.RUNNING`

	__GongEnumValueEntry__000020_SLEEP_100_MS.Name = `SLEEP_100_MS`

	//gong:ident [ref_models.EngineDriverState.SLEEP_100_MS] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000020_SLEEP_100_MS.Identifier = `ref_models.EngineDriverState.SLEEP_100_MS`

	__GongEnumValueEntry__000021_STATE_CHANGED.Name = `STATE_CHANGED`

	//gong:ident [ref_models.EngineStopMode.STATE_CHANGED] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000021_STATE_CHANGED.Identifier = `ref_models.EngineStopMode.STATE_CHANGED`

	__GongEnumValueEntry__000022_TEN_MINUTES.Name = `TEN_MINUTES`

	//gong:ident [ref_models.EngineStopMode.TEN_MINUTES] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000022_TEN_MINUTES.Identifier = `ref_models.EngineStopMode.TEN_MINUTES`

	__GongEnumValueEntry__000023_UNKOWN.Name = `UNKOWN`

	//gong:ident [ref_models.EngineDriverState.UNKOWN] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000023_UNKOWN.Identifier = `ref_models.EngineDriverState.UNKOWN`

	__GongStructShape__000000_NewDiagram_DummyAgent.Name = `NewDiagram-DummyAgent`

	//gong:ident [ref_models.DummyAgent] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_NewDiagram_DummyAgent.Identifier = `ref_models.DummyAgent`
	__GongStructShape__000000_NewDiagram_DummyAgent.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_DummyAgent.NbInstances = 0
	__GongStructShape__000000_NewDiagram_DummyAgent.Width = 240.000000
	__GongStructShape__000000_NewDiagram_DummyAgent.Height = 93.000000
	__GongStructShape__000000_NewDiagram_DummyAgent.IsSelected = false

	__GongStructShape__000001_NewDiagram_Engine.Name = `NewDiagram-Engine`

	//gong:ident [ref_models.Engine] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_NewDiagram_Engine.Identifier = `ref_models.Engine`
	__GongStructShape__000001_NewDiagram_Engine.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Engine.NbInstances = 1
	__GongStructShape__000001_NewDiagram_Engine.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Engine.Height = 183.000000
	__GongStructShape__000001_NewDiagram_Engine.IsSelected = false

	__GongStructShape__000002_NewDiagram_Event.Name = `NewDiagram-Event`

	//gong:ident [ref_models.Event] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_NewDiagram_Event.Identifier = `ref_models.Event`
	__GongStructShape__000002_NewDiagram_Event.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_Event.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Event.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Event.Height = 93.000000
	__GongStructShape__000002_NewDiagram_Event.IsSelected = false

	__GongStructShape__000003_NewDiagram_GongsimCommand.Name = `NewDiagram-Command`

	//gong:ident [ref_models.Command] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_NewDiagram_GongsimCommand.Identifier = `ref_models.Command`
	__GongStructShape__000003_NewDiagram_GongsimCommand.ShowNbInstances = false
	__GongStructShape__000003_NewDiagram_GongsimCommand.NbInstances = 1
	__GongStructShape__000003_NewDiagram_GongsimCommand.Width = 389.000000
	__GongStructShape__000003_NewDiagram_GongsimCommand.Height = 138.000000
	__GongStructShape__000003_NewDiagram_GongsimCommand.IsSelected = false

	__GongStructShape__000004_NewDiagram_Status.Name = `NewDiagram-Status`

	//gong:ident [ref_models.Status] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_NewDiagram_Status.Identifier = `ref_models.Status`
	__GongStructShape__000004_NewDiagram_Status.ShowNbInstances = false
	__GongStructShape__000004_NewDiagram_Status.NbInstances = 1
	__GongStructShape__000004_NewDiagram_Status.Width = 365.000000
	__GongStructShape__000004_NewDiagram_Status.Height = 138.000000
	__GongStructShape__000004_NewDiagram_Status.IsSelected = false

	__Link__000000_Engine.Name = `Engine`

	//gong:ident [ref_models.Command.Engine] comment added to overcome the problem with the comment map association
	__Link__000000_Engine.Identifier = `ref_models.Command.Engine`

	//gong:ident [ref_models.Engine] comment added to overcome the problem with the comment map association
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

	__Position__000000_Pos_NewDiagram_ControlMode.X = 798.000000
	__Position__000000_Pos_NewDiagram_ControlMode.Y = 783.000000
	__Position__000000_Pos_NewDiagram_ControlMode.Name = `Pos-NewDiagram-ControlMode`

	__Position__000001_Pos_NewDiagram_DummyAgent.X = 50.000000
	__Position__000001_Pos_NewDiagram_DummyAgent.Y = 670.000000
	__Position__000001_Pos_NewDiagram_DummyAgent.Name = `Pos-NewDiagram-DummyAgent`

	__Position__000002_Pos_NewDiagram_Engine.X = 44.000000
	__Position__000002_Pos_NewDiagram_Engine.Y = 85.000000
	__Position__000002_Pos_NewDiagram_Engine.Name = `Pos-NewDiagram-Engine`

	__Position__000003_Pos_NewDiagram_EngineDriverState.X = 785.000000
	__Position__000003_Pos_NewDiagram_EngineDriverState.Y = 589.000000
	__Position__000003_Pos_NewDiagram_EngineDriverState.Name = `Pos-NewDiagram-EngineDriverState`

	__Position__000004_Pos_NewDiagram_EngineRunMode.X = 476.000000
	__Position__000004_Pos_NewDiagram_EngineRunMode.Y = 259.000000
	__Position__000004_Pos_NewDiagram_EngineRunMode.Name = `Pos-NewDiagram-EngineRunMode`

	__Position__000005_Pos_NewDiagram_EngineState.X = 488.000000
	__Position__000005_Pos_NewDiagram_EngineState.Y = 86.000000
	__Position__000005_Pos_NewDiagram_EngineState.Name = `Pos-NewDiagram-EngineState`

	__Position__000006_Pos_NewDiagram_EngineStopMode.X = 780.000000
	__Position__000006_Pos_NewDiagram_EngineStopMode.Y = 431.000000
	__Position__000006_Pos_NewDiagram_EngineStopMode.Name = `Pos-NewDiagram-EngineStopMode`

	__Position__000007_Pos_NewDiagram_Event.X = 60.000000
	__Position__000007_Pos_NewDiagram_Event.Y = 520.000000
	__Position__000007_Pos_NewDiagram_Event.Name = `Pos-NewDiagram-Event`

	__Position__000008_Pos_NewDiagram_GongsimCommand.X = 29.000000
	__Position__000008_Pos_NewDiagram_GongsimCommand.Y = 318.000000
	__Position__000008_Pos_NewDiagram_GongsimCommand.Name = `Pos-NewDiagram-Command`

	__Position__000009_Pos_NewDiagram_GongsimCommandType.X = 778.000000
	__Position__000009_Pos_NewDiagram_GongsimCommandType.Y = 236.000000
	__Position__000009_Pos_NewDiagram_GongsimCommandType.Name = `Pos-NewDiagram-GongsimCommandType`

	__Position__000010_Pos_NewDiagram_Status.X = 367.000000
	__Position__000010_Pos_NewDiagram_Status.Y = 534.000000
	__Position__000010_Pos_NewDiagram_Status.Name = `Pos-NewDiagram-Status`

	__Position__000011_Pos_NewDiagram_SpeedCommandType.X = 777.000000
	__Position__000011_Pos_NewDiagram_SpeedCommandType.Y = 83.000000
	__Position__000011_Pos_NewDiagram_SpeedCommandType.Name = `Pos-NewDiagram-SpeedCommandType`

	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine.X = 644.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine.Y = 274.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Command and NewDiagram-Engine`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Engine)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_GongsimCommand)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_Status)
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
	__GongStructShape__000000_NewDiagram_DummyAgent.Fields = append(__GongStructShape__000000_NewDiagram_DummyAgent.Fields, __Field__000011_Name)
	__GongStructShape__000001_NewDiagram_Engine.Position = __Position__000002_Pos_NewDiagram_Engine
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000012_Name)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000009_EndTime)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000006_CurrentTime)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000016_SecondsSinceStart)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000010_Fired)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000003_ControlMode)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000020_State)
	__GongStructShape__000001_NewDiagram_Engine.Fields = append(__GongStructShape__000001_NewDiagram_Engine.Fields, __Field__000017_Speed)
	__GongStructShape__000002_NewDiagram_Event.Position = __Position__000007_Pos_NewDiagram_Event
	__GongStructShape__000002_NewDiagram_Event.Fields = append(__GongStructShape__000002_NewDiagram_Event.Fields, __Field__000014_Name)
	__GongStructShape__000002_NewDiagram_Event.Fields = append(__GongStructShape__000002_NewDiagram_Event.Fields, __Field__000008_Duration)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Position = __Position__000008_Pos_NewDiagram_GongsimCommand
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000015_Name)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000000_Command)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000001_CommandDate)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000019_SpeedCommandType)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Fields, __Field__000007_DateSpeedCommand)
	__GongStructShape__000003_NewDiagram_GongsimCommand.Links = append(__GongStructShape__000003_NewDiagram_GongsimCommand.Links, __Link__000000_Engine)
	__GongStructShape__000004_NewDiagram_Status.Position = __Position__000010_Pos_NewDiagram_Status
	__GongStructShape__000004_NewDiagram_Status.Fields = append(__GongStructShape__000004_NewDiagram_Status.Fields, __Field__000013_Name)
	__GongStructShape__000004_NewDiagram_Status.Fields = append(__GongStructShape__000004_NewDiagram_Status.Fields, __Field__000004_CurrentCommand)
	__GongStructShape__000004_NewDiagram_Status.Fields = append(__GongStructShape__000004_NewDiagram_Status.Fields, __Field__000002_CompletionDate)
	__GongStructShape__000004_NewDiagram_Status.Fields = append(__GongStructShape__000004_NewDiagram_Status.Fields, __Field__000005_CurrentSpeedCommand)
	__GongStructShape__000004_NewDiagram_Status.Fields = append(__GongStructShape__000004_NewDiagram_Status.Fields, __Field__000018_SpeedCommandCompletionDate)
	__Link__000000_Engine.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_GongsimCommand_and_NewDiagram_Engine
}
