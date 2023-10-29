package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongfly/go/models"
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

	"ref_models.Aircraft_": ref_models.Aircraft_,

	"ref_models.Center_": ref_models.Center_,

	"ref_models.CivilianAirport": &(ref_models.CivilianAirport{}),

	"ref_models.CivilianAirport.Lat": (ref_models.CivilianAirport{}).Lat,

	"ref_models.CivilianAirport.Lng": (ref_models.CivilianAirport{}).Lng,

	"ref_models.CivilianAirport.Name": (ref_models.CivilianAirport{}).Name,

	"ref_models.ConceptEnum": ref_models.ConceptEnum(""),

	"ref_models.EN_ROUTE_NOMINAL": ref_models.EN_ROUTE_NOMINAL,

	"ref_models.LANDED": ref_models.LANDED,

	"ref_models.Liner": &(ref_models.Liner{}),

	"ref_models.Liner.DistanceToTarget": (ref_models.Liner{}).DistanceToTarget,

	"ref_models.Liner.Heading": (ref_models.Liner{}).Heading,

	"ref_models.Liner.Lat": (ref_models.Liner{}).Lat,

	"ref_models.Liner.Level": (ref_models.Liner{}).Level,

	"ref_models.Liner.Lng": (ref_models.Liner{}).Lng,

	"ref_models.Liner.MaxRotationalSpeed": (ref_models.Liner{}).MaxRotationalSpeed,

	"ref_models.Liner.Name": (ref_models.Liner{}).Name,

	"ref_models.Liner.ReporingLine": (ref_models.Liner{}).ReporingLine,

	"ref_models.Liner.Speed": (ref_models.Liner{}).Speed,

	"ref_models.Liner.State": (ref_models.Liner{}).State,

	"ref_models.Liner.TargetHeading": (ref_models.Liner{}).TargetHeading,

	"ref_models.Liner.TargetLocationLat": (ref_models.Liner{}).TargetLocationLat,

	"ref_models.Liner.TargetLocationLng": (ref_models.Liner{}).TargetLocationLng,

	"ref_models.Liner.Timestampstring": (ref_models.Liner{}).Timestampstring,

	"ref_models.Liner.VerticalSpeed": (ref_models.Liner{}).VerticalSpeed,

	"ref_models.LinerStateEnum": ref_models.LinerStateEnum(""),

	"ref_models.MESSAGE_ARRIVED": ref_models.MESSAGE_ARRIVED,

	"ref_models.MESSAGE_EN_ROUTE": ref_models.MESSAGE_EN_ROUTE,

	"ref_models.Message": &(ref_models.Message{}),

	"ref_models.Message.About_string": (ref_models.Message{}).About_string,

	"ref_models.Message.Content": (ref_models.Message{}).Content,

	"ref_models.Message.Destination": (ref_models.Message{}).Destination,

	"ref_models.Message.Display": (ref_models.Message{}).Display,

	"ref_models.Message.DistanceToTarget": (ref_models.Message{}).DistanceToTarget,

	"ref_models.Message.DurationSinceSimulationStart": (ref_models.Message{}).DurationSinceSimulationStart,

	"ref_models.Message.Heading": (ref_models.Message{}).Heading,

	"ref_models.Message.Lat": (ref_models.Message{}).Lat,

	"ref_models.Message.Level": (ref_models.Message{}).Level,

	"ref_models.Message.Lng": (ref_models.Message{}).Lng,

	"ref_models.Message.Name": (ref_models.Message{}).Name,

	"ref_models.Message.Source": (ref_models.Message{}).Source,

	"ref_models.Message.Speed": (ref_models.Message{}).Speed,

	"ref_models.Message.State": (ref_models.Message{}).State,

	"ref_models.Message.TargetLocationLat": (ref_models.Message{}).TargetLocationLat,

	"ref_models.Message.TargetLocationLng": (ref_models.Message{}).TargetLocationLng,

	"ref_models.Message.Timestampstartstring": (ref_models.Message{}).Timestampstartstring,

	"ref_models.Message.Timestampstring": (ref_models.Message{}).Timestampstring,

	"ref_models.MessageStateEnum": ref_models.MessageStateEnum(""),

	"ref_models.Network_": ref_models.Network_,

	"ref_models.OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING": ref_models.OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING,

	"ref_models.OPS_COM_LINK_OPERATIONAL_LINE_WORKING": ref_models.OPS_COM_LINK_OPERATIONAL_LINE_WORKING,

	"ref_models.OperationalLineStateEnum": ref_models.OperationalLineStateEnum(""),

	"ref_models.OpsLine": &(ref_models.OpsLine{}),

	"ref_models.OpsLine.IsTransmitting": (ref_models.OpsLine{}).IsTransmitting,

	"ref_models.OpsLine.IsTransmittingBackward": (ref_models.OpsLine{}).IsTransmittingBackward,

	"ref_models.OpsLine.Name": (ref_models.OpsLine{}).Name,

	"ref_models.OpsLine.Scenario": (ref_models.OpsLine{}).Scenario,

	"ref_models.OpsLine.State": (ref_models.OpsLine{}).State,

	"ref_models.OpsLine.TransmissionMessage": (ref_models.OpsLine{}).TransmissionMessage,

	"ref_models.OpsLine.TransmissionMessageBackward": (ref_models.OpsLine{}).TransmissionMessageBackward,

	"ref_models.OrderEnum": ref_models.OrderEnum(""),

	"ref_models.Radar": &(ref_models.Radar{}),

	"ref_models.Radar.Lat": (ref_models.Radar{}).Lat,

	"ref_models.Radar.Lng": (ref_models.Radar{}).Lng,

	"ref_models.Radar.Name": (ref_models.Radar{}).Name,

	"ref_models.Radar.Range": (ref_models.Radar{}).Range,

	"ref_models.Radar.State": (ref_models.Radar{}).State,

	"ref_models.RadarStateEnum": ref_models.RadarStateEnum(""),

	"ref_models.ReportEnum": ref_models.ReportEnum(""),

	"ref_models.Satellite": &(ref_models.Satellite{}),

	"ref_models.Satellite.Heading": (ref_models.Satellite{}).Heading,

	"ref_models.Satellite.Lat": (ref_models.Satellite{}).Lat,

	"ref_models.Satellite.Level": (ref_models.Satellite{}).Level,

	"ref_models.Satellite.Line1": (ref_models.Satellite{}).Line1,

	"ref_models.Satellite.Line2": (ref_models.Satellite{}).Line2,

	"ref_models.Satellite.Lng": (ref_models.Satellite{}).Lng,

	"ref_models.Satellite.Name": (ref_models.Satellite{}).Name,

	"ref_models.Satellite.Speed": (ref_models.Satellite{}).Speed,

	"ref_models.Satellite.Timestampstring": (ref_models.Satellite{}).Timestampstring,

	"ref_models.Satellite.VerticalSpeed": (ref_models.Satellite{}).VerticalSpeed,

	"ref_models.Satellite_": ref_models.Satellite_,

	"ref_models.Scenario": &(ref_models.Scenario{}),

	"ref_models.Scenario.Lat": (ref_models.Scenario{}).Lat,

	"ref_models.Scenario.Lng": (ref_models.Scenario{}).Lng,

	"ref_models.Scenario.MessageVisualSpeed": (ref_models.Scenario{}).MessageVisualSpeed,

	"ref_models.Scenario.Name": (ref_models.Scenario{}).Name,

	"ref_models.Scenario.ZoomLevel": (ref_models.Scenario{}).ZoomLevel,

	"ref_models.System_": ref_models.System_,

	"ref_models.TAKE_OFF": ref_models.TAKE_OFF,

	"ref_models.TAKE_OFF_COMPLETED": ref_models.TAKE_OFF_COMPLETED,

	"ref_models.WORKING": ref_models.WORKING,
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

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_About_string := (&models.Field{Name: `About_string`}).Stage(stage)
	__Field__000001_Content := (&models.Field{Name: `Content`}).Stage(stage)
	__Field__000002_Destination := (&models.Field{Name: `Destination`}).Stage(stage)
	__Field__000003_Display := (&models.Field{Name: `Display`}).Stage(stage)
	__Field__000004_DistanceToTarget := (&models.Field{Name: `DistanceToTarget`}).Stage(stage)
	__Field__000005_DistanceToTarget := (&models.Field{Name: `DistanceToTarget`}).Stage(stage)
	__Field__000006_DurationSinceSimulationStart := (&models.Field{Name: `DurationSinceSimulationStart`}).Stage(stage)
	__Field__000007_Heading := (&models.Field{Name: `Heading`}).Stage(stage)
	__Field__000008_Heading := (&models.Field{Name: `Heading`}).Stage(stage)
	__Field__000009_Heading := (&models.Field{Name: `Heading`}).Stage(stage)
	__Field__000010_IsTransmitting := (&models.Field{Name: `IsTransmitting`}).Stage(stage)
	__Field__000011_IsTransmittingBackward := (&models.Field{Name: `IsTransmittingBackward`}).Stage(stage)
	__Field__000012_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000013_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000014_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000015_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000016_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000017_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000018_Level := (&models.Field{Name: `Level`}).Stage(stage)
	__Field__000019_Level := (&models.Field{Name: `Level`}).Stage(stage)
	__Field__000020_Level := (&models.Field{Name: `Level`}).Stage(stage)
	__Field__000021_Line1 := (&models.Field{Name: `Line1`}).Stage(stage)
	__Field__000022_Line2 := (&models.Field{Name: `Line2`}).Stage(stage)
	__Field__000023_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000024_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000025_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000026_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000027_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000028_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000029_MaxRotationalSpeed := (&models.Field{Name: `MaxRotationalSpeed`}).Stage(stage)
	__Field__000030_MessageVisualSpeed := (&models.Field{Name: `MessageVisualSpeed`}).Stage(stage)
	__Field__000031_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000032_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000033_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000034_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000035_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000036_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000037_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000038_Range := (&models.Field{Name: `Range`}).Stage(stage)
	__Field__000039_Source := (&models.Field{Name: `Source`}).Stage(stage)
	__Field__000040_Speed := (&models.Field{Name: `Speed`}).Stage(stage)
	__Field__000041_Speed := (&models.Field{Name: `Speed`}).Stage(stage)
	__Field__000042_Speed := (&models.Field{Name: `Speed`}).Stage(stage)
	__Field__000043_State := (&models.Field{Name: `State`}).Stage(stage)
	__Field__000044_State := (&models.Field{Name: `State`}).Stage(stage)
	__Field__000045_State := (&models.Field{Name: `State`}).Stage(stage)
	__Field__000046_State := (&models.Field{Name: `State`}).Stage(stage)
	__Field__000047_TargetHeading := (&models.Field{Name: `TargetHeading`}).Stage(stage)
	__Field__000048_TargetLocationLat := (&models.Field{Name: `TargetLocationLat`}).Stage(stage)
	__Field__000049_TargetLocationLat := (&models.Field{Name: `TargetLocationLat`}).Stage(stage)
	__Field__000050_TargetLocationLng := (&models.Field{Name: `TargetLocationLng`}).Stage(stage)
	__Field__000051_TargetLocationLng := (&models.Field{Name: `TargetLocationLng`}).Stage(stage)
	__Field__000052_Timestampstartstring := (&models.Field{Name: `Timestampstartstring`}).Stage(stage)
	__Field__000053_Timestampstring := (&models.Field{Name: `Timestampstring`}).Stage(stage)
	__Field__000054_Timestampstring := (&models.Field{Name: `Timestampstring`}).Stage(stage)
	__Field__000055_Timestampstring := (&models.Field{Name: `Timestampstring`}).Stage(stage)
	__Field__000056_TransmissionMessage := (&models.Field{Name: `TransmissionMessage`}).Stage(stage)
	__Field__000057_TransmissionMessageBackward := (&models.Field{Name: `TransmissionMessageBackward`}).Stage(stage)
	__Field__000058_VerticalSpeed := (&models.Field{Name: `VerticalSpeed`}).Stage(stage)
	__Field__000059_VerticalSpeed := (&models.Field{Name: `VerticalSpeed`}).Stage(stage)
	__Field__000060_ZoomLevel := (&models.Field{Name: `ZoomLevel`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_ConceptEnum := (&models.GongEnumShape{Name: `NewDiagram-ConceptEnum`}).Stage(stage)
	__GongEnumShape__000001_NewDiagram_LinerStateEnum := (&models.GongEnumShape{Name: `NewDiagram-LinerStateEnum`}).Stage(stage)
	__GongEnumShape__000002_NewDiagram_MessageStateEnum := (&models.GongEnumShape{Name: `NewDiagram-MessageStateEnum`}).Stage(stage)
	__GongEnumShape__000003_NewDiagram_OperationalLineStateEnum := (&models.GongEnumShape{Name: `NewDiagram-OperationalLineStateEnum`}).Stage(stage)
	__GongEnumShape__000004_NewDiagram_OrderEnum := (&models.GongEnumShape{Name: `NewDiagram-OrderEnum`}).Stage(stage)
	__GongEnumShape__000005_NewDiagram_RadarStateEnum := (&models.GongEnumShape{Name: `NewDiagram-RadarStateEnum`}).Stage(stage)
	__GongEnumShape__000006_NewDiagram_ReportEnum := (&models.GongEnumShape{Name: `NewDiagram-ReportEnum`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_CivilianAirport := (&models.GongStructShape{Name: `NewDiagram-CivilianAirport`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Liner := (&models.GongStructShape{Name: `NewDiagram-Liner`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_Message := (&models.GongStructShape{Name: `NewDiagram-Message`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_OpsLine := (&models.GongStructShape{Name: `NewDiagram-OpsLine`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_Radar := (&models.GongStructShape{Name: `NewDiagram-Radar`}).Stage(stage)
	__GongStructShape__000005_NewDiagram_Satellite := (&models.GongStructShape{Name: `NewDiagram-Satellite`}).Stage(stage)
	__GongStructShape__000006_NewDiagram_Scenario := (&models.GongStructShape{Name: `NewDiagram-Scenario`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_ReporingLine := (&models.Link{Name: `ReporingLine`}).Stage(stage)
	__Link__000001_Scenario := (&models.Link{Name: `Scenario`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_CivilianAirport := (&models.Position{Name: `Pos-NewDiagram-CivilianAirport`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_ConceptEnum := (&models.Position{Name: `Pos-NewDiagram-ConceptEnum`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Liner := (&models.Position{Name: `Pos-NewDiagram-Liner`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_LinerStateEnum := (&models.Position{Name: `Pos-NewDiagram-LinerStateEnum`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_Message := (&models.Position{Name: `Pos-NewDiagram-Message`}).Stage(stage)
	__Position__000005_Pos_NewDiagram_MessageStateEnum := (&models.Position{Name: `Pos-NewDiagram-MessageStateEnum`}).Stage(stage)
	__Position__000006_Pos_NewDiagram_OperationalLineStateEnum := (&models.Position{Name: `Pos-NewDiagram-OperationalLineStateEnum`}).Stage(stage)
	__Position__000007_Pos_NewDiagram_OpsLine := (&models.Position{Name: `Pos-NewDiagram-OpsLine`}).Stage(stage)
	__Position__000008_Pos_NewDiagram_OrderEnum := (&models.Position{Name: `Pos-NewDiagram-OrderEnum`}).Stage(stage)
	__Position__000009_Pos_NewDiagram_Radar := (&models.Position{Name: `Pos-NewDiagram-Radar`}).Stage(stage)
	__Position__000010_Pos_NewDiagram_RadarStateEnum := (&models.Position{Name: `Pos-NewDiagram-RadarStateEnum`}).Stage(stage)
	__Position__000011_Pos_NewDiagram_ReportEnum := (&models.Position{Name: `Pos-NewDiagram-ReportEnum`}).Stage(stage)
	__Position__000012_Pos_NewDiagram_Satellite := (&models.Position{Name: `Pos-NewDiagram-Satellite`}).Stage(stage)
	__Position__000013_Pos_NewDiagram_Scenario := (&models.Position{Name: `Pos-NewDiagram-Scenario`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Liner and NewDiagram-OpsLine`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_OpsLine_and_NewDiagram_Scenario := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-OpsLine and NewDiagram-Scenario`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_About_string.Name = `About_string`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.About_string]
	__Field__000000_About_string.Identifier = `ref_models.Message.About_string`
	__Field__000000_About_string.FieldTypeAsString = ``
	__Field__000000_About_string.Structname = `Message`
	__Field__000000_About_string.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Content.Name = `Content`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Content]
	__Field__000001_Content.Identifier = `ref_models.Message.Content`
	__Field__000001_Content.FieldTypeAsString = ``
	__Field__000001_Content.Structname = `Message`
	__Field__000001_Content.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Destination.Name = `Destination`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Destination]
	__Field__000002_Destination.Identifier = `ref_models.Message.Destination`
	__Field__000002_Destination.FieldTypeAsString = ``
	__Field__000002_Destination.Structname = `Message`
	__Field__000002_Destination.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Display.Name = `Display`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Display]
	__Field__000003_Display.Identifier = `ref_models.Message.Display`
	__Field__000003_Display.FieldTypeAsString = ``
	__Field__000003_Display.Structname = `Message`
	__Field__000003_Display.Fieldtypename = `bool`

	// Field values setup
	__Field__000004_DistanceToTarget.Name = `DistanceToTarget`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.DistanceToTarget]
	__Field__000004_DistanceToTarget.Identifier = `ref_models.Liner.DistanceToTarget`
	__Field__000004_DistanceToTarget.FieldTypeAsString = ``
	__Field__000004_DistanceToTarget.Structname = `Liner`
	__Field__000004_DistanceToTarget.Fieldtypename = `float64`

	// Field values setup
	__Field__000005_DistanceToTarget.Name = `DistanceToTarget`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.DistanceToTarget]
	__Field__000005_DistanceToTarget.Identifier = `ref_models.Message.DistanceToTarget`
	__Field__000005_DistanceToTarget.FieldTypeAsString = ``
	__Field__000005_DistanceToTarget.Structname = `Message`
	__Field__000005_DistanceToTarget.Fieldtypename = `float64`

	// Field values setup
	__Field__000006_DurationSinceSimulationStart.Name = `DurationSinceSimulationStart`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.DurationSinceSimulationStart]
	__Field__000006_DurationSinceSimulationStart.Identifier = `ref_models.Message.DurationSinceSimulationStart`
	__Field__000006_DurationSinceSimulationStart.FieldTypeAsString = ``
	__Field__000006_DurationSinceSimulationStart.Structname = `Message`
	__Field__000006_DurationSinceSimulationStart.Fieldtypename = `Duration`

	// Field values setup
	__Field__000007_Heading.Name = `Heading`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Heading]
	__Field__000007_Heading.Identifier = `ref_models.Liner.Heading`
	__Field__000007_Heading.FieldTypeAsString = ``
	__Field__000007_Heading.Structname = `Liner`
	__Field__000007_Heading.Fieldtypename = `float64`

	// Field values setup
	__Field__000008_Heading.Name = `Heading`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Heading]
	__Field__000008_Heading.Identifier = `ref_models.Message.Heading`
	__Field__000008_Heading.FieldTypeAsString = ``
	__Field__000008_Heading.Structname = `Message`
	__Field__000008_Heading.Fieldtypename = `float64`

	// Field values setup
	__Field__000009_Heading.Name = `Heading`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Heading]
	__Field__000009_Heading.Identifier = `ref_models.Satellite.Heading`
	__Field__000009_Heading.FieldTypeAsString = ``
	__Field__000009_Heading.Structname = `Satellite`
	__Field__000009_Heading.Fieldtypename = `float64`

	// Field values setup
	__Field__000010_IsTransmitting.Name = `IsTransmitting`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.IsTransmitting]
	__Field__000010_IsTransmitting.Identifier = `ref_models.OpsLine.IsTransmitting`
	__Field__000010_IsTransmitting.FieldTypeAsString = ``
	__Field__000010_IsTransmitting.Structname = `OpsLine`
	__Field__000010_IsTransmitting.Fieldtypename = `bool`

	// Field values setup
	__Field__000011_IsTransmittingBackward.Name = `IsTransmittingBackward`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.IsTransmittingBackward]
	__Field__000011_IsTransmittingBackward.Identifier = `ref_models.OpsLine.IsTransmittingBackward`
	__Field__000011_IsTransmittingBackward.FieldTypeAsString = ``
	__Field__000011_IsTransmittingBackward.Structname = `OpsLine`
	__Field__000011_IsTransmittingBackward.Fieldtypename = `bool`

	// Field values setup
	__Field__000012_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Lat]
	__Field__000012_Lat.Identifier = `ref_models.Satellite.Lat`
	__Field__000012_Lat.FieldTypeAsString = ``
	__Field__000012_Lat.Structname = `Satellite`
	__Field__000012_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000013_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Lat]
	__Field__000013_Lat.Identifier = `ref_models.Liner.Lat`
	__Field__000013_Lat.FieldTypeAsString = ``
	__Field__000013_Lat.Structname = `Liner`
	__Field__000013_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000014_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Radar.Lat]
	__Field__000014_Lat.Identifier = `ref_models.Radar.Lat`
	__Field__000014_Lat.FieldTypeAsString = ``
	__Field__000014_Lat.Structname = `Radar`
	__Field__000014_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000015_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport.Lat]
	__Field__000015_Lat.Identifier = `ref_models.CivilianAirport.Lat`
	__Field__000015_Lat.FieldTypeAsString = ``
	__Field__000015_Lat.Structname = `CivilianAirport`
	__Field__000015_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000016_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Lat]
	__Field__000016_Lat.Identifier = `ref_models.Message.Lat`
	__Field__000016_Lat.FieldTypeAsString = ``
	__Field__000016_Lat.Structname = `Message`
	__Field__000016_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000017_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario.Lat]
	__Field__000017_Lat.Identifier = `ref_models.Scenario.Lat`
	__Field__000017_Lat.FieldTypeAsString = ``
	__Field__000017_Lat.Structname = `Scenario`
	__Field__000017_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000018_Level.Name = `Level`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Level]
	__Field__000018_Level.Identifier = `ref_models.Satellite.Level`
	__Field__000018_Level.FieldTypeAsString = ``
	__Field__000018_Level.Structname = `Satellite`
	__Field__000018_Level.Fieldtypename = `float64`

	// Field values setup
	__Field__000019_Level.Name = `Level`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Level]
	__Field__000019_Level.Identifier = `ref_models.Message.Level`
	__Field__000019_Level.FieldTypeAsString = ``
	__Field__000019_Level.Structname = `Message`
	__Field__000019_Level.Fieldtypename = `float64`

	// Field values setup
	__Field__000020_Level.Name = `Level`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Level]
	__Field__000020_Level.Identifier = `ref_models.Liner.Level`
	__Field__000020_Level.FieldTypeAsString = ``
	__Field__000020_Level.Structname = `Liner`
	__Field__000020_Level.Fieldtypename = `float64`

	// Field values setup
	__Field__000021_Line1.Name = `Line1`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Line1]
	__Field__000021_Line1.Identifier = `ref_models.Satellite.Line1`
	__Field__000021_Line1.FieldTypeAsString = ``
	__Field__000021_Line1.Structname = `Satellite`
	__Field__000021_Line1.Fieldtypename = `string`

	// Field values setup
	__Field__000022_Line2.Name = `Line2`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Line2]
	__Field__000022_Line2.Identifier = `ref_models.Satellite.Line2`
	__Field__000022_Line2.FieldTypeAsString = ``
	__Field__000022_Line2.Structname = `Satellite`
	__Field__000022_Line2.Fieldtypename = `string`

	// Field values setup
	__Field__000023_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport.Lng]
	__Field__000023_Lng.Identifier = `ref_models.CivilianAirport.Lng`
	__Field__000023_Lng.FieldTypeAsString = ``
	__Field__000023_Lng.Structname = `CivilianAirport`
	__Field__000023_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000024_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Lng]
	__Field__000024_Lng.Identifier = `ref_models.Message.Lng`
	__Field__000024_Lng.FieldTypeAsString = ``
	__Field__000024_Lng.Structname = `Message`
	__Field__000024_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000025_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Lng]
	__Field__000025_Lng.Identifier = `ref_models.Liner.Lng`
	__Field__000025_Lng.FieldTypeAsString = ``
	__Field__000025_Lng.Structname = `Liner`
	__Field__000025_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000026_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Radar.Lng]
	__Field__000026_Lng.Identifier = `ref_models.Radar.Lng`
	__Field__000026_Lng.FieldTypeAsString = ``
	__Field__000026_Lng.Structname = `Radar`
	__Field__000026_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000027_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario.Lng]
	__Field__000027_Lng.Identifier = `ref_models.Scenario.Lng`
	__Field__000027_Lng.FieldTypeAsString = ``
	__Field__000027_Lng.Structname = `Scenario`
	__Field__000027_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000028_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Lng]
	__Field__000028_Lng.Identifier = `ref_models.Satellite.Lng`
	__Field__000028_Lng.FieldTypeAsString = ``
	__Field__000028_Lng.Structname = `Satellite`
	__Field__000028_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000029_MaxRotationalSpeed.Name = `MaxRotationalSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.MaxRotationalSpeed]
	__Field__000029_MaxRotationalSpeed.Identifier = `ref_models.Liner.MaxRotationalSpeed`
	__Field__000029_MaxRotationalSpeed.FieldTypeAsString = ``
	__Field__000029_MaxRotationalSpeed.Structname = `Liner`
	__Field__000029_MaxRotationalSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000030_MessageVisualSpeed.Name = `MessageVisualSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario.MessageVisualSpeed]
	__Field__000030_MessageVisualSpeed.Identifier = `ref_models.Scenario.MessageVisualSpeed`
	__Field__000030_MessageVisualSpeed.FieldTypeAsString = ``
	__Field__000030_MessageVisualSpeed.Structname = `Scenario`
	__Field__000030_MessageVisualSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000031_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Name]
	__Field__000031_Name.Identifier = `ref_models.Message.Name`
	__Field__000031_Name.FieldTypeAsString = ``
	__Field__000031_Name.Structname = `Message`
	__Field__000031_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000032_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Name]
	__Field__000032_Name.Identifier = `ref_models.Satellite.Name`
	__Field__000032_Name.FieldTypeAsString = ``
	__Field__000032_Name.Structname = `Satellite`
	__Field__000032_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000033_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.Name]
	__Field__000033_Name.Identifier = `ref_models.OpsLine.Name`
	__Field__000033_Name.FieldTypeAsString = ``
	__Field__000033_Name.Structname = `OpsLine`
	__Field__000033_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000034_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport.Name]
	__Field__000034_Name.Identifier = `ref_models.CivilianAirport.Name`
	__Field__000034_Name.FieldTypeAsString = ``
	__Field__000034_Name.Structname = `CivilianAirport`
	__Field__000034_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000035_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Name]
	__Field__000035_Name.Identifier = `ref_models.Liner.Name`
	__Field__000035_Name.FieldTypeAsString = ``
	__Field__000035_Name.Structname = `Liner`
	__Field__000035_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000036_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Radar.Name]
	__Field__000036_Name.Identifier = `ref_models.Radar.Name`
	__Field__000036_Name.FieldTypeAsString = ``
	__Field__000036_Name.Structname = `Radar`
	__Field__000036_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000037_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario.Name]
	__Field__000037_Name.Identifier = `ref_models.Scenario.Name`
	__Field__000037_Name.FieldTypeAsString = ``
	__Field__000037_Name.Structname = `Scenario`
	__Field__000037_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000038_Range.Name = `Range`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Radar.Range]
	__Field__000038_Range.Identifier = `ref_models.Radar.Range`
	__Field__000038_Range.FieldTypeAsString = ``
	__Field__000038_Range.Structname = `Radar`
	__Field__000038_Range.Fieldtypename = `float64`

	// Field values setup
	__Field__000039_Source.Name = `Source`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Source]
	__Field__000039_Source.Identifier = `ref_models.Message.Source`
	__Field__000039_Source.FieldTypeAsString = ``
	__Field__000039_Source.Structname = `Message`
	__Field__000039_Source.Fieldtypename = `string`

	// Field values setup
	__Field__000040_Speed.Name = `Speed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Speed]
	__Field__000040_Speed.Identifier = `ref_models.Message.Speed`
	__Field__000040_Speed.FieldTypeAsString = ``
	__Field__000040_Speed.Structname = `Message`
	__Field__000040_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000041_Speed.Name = `Speed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Speed]
	__Field__000041_Speed.Identifier = `ref_models.Liner.Speed`
	__Field__000041_Speed.FieldTypeAsString = ``
	__Field__000041_Speed.Structname = `Liner`
	__Field__000041_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000042_Speed.Name = `Speed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Speed]
	__Field__000042_Speed.Identifier = `ref_models.Satellite.Speed`
	__Field__000042_Speed.FieldTypeAsString = ``
	__Field__000042_Speed.Structname = `Satellite`
	__Field__000042_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000043_State.Name = `State`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Radar.State]
	__Field__000043_State.Identifier = `ref_models.Radar.State`
	__Field__000043_State.FieldTypeAsString = ``
	__Field__000043_State.Structname = `Radar`
	__Field__000043_State.Fieldtypename = `RadarStateEnum`

	// Field values setup
	__Field__000044_State.Name = `State`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.State]
	__Field__000044_State.Identifier = `ref_models.Message.State`
	__Field__000044_State.FieldTypeAsString = ``
	__Field__000044_State.Structname = `Message`
	__Field__000044_State.Fieldtypename = `MessageStateEnum`

	// Field values setup
	__Field__000045_State.Name = `State`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.State]
	__Field__000045_State.Identifier = `ref_models.Liner.State`
	__Field__000045_State.FieldTypeAsString = ``
	__Field__000045_State.Structname = `Liner`
	__Field__000045_State.Fieldtypename = `LinerStateEnum`

	// Field values setup
	__Field__000046_State.Name = `State`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.State]
	__Field__000046_State.Identifier = `ref_models.OpsLine.State`
	__Field__000046_State.FieldTypeAsString = ``
	__Field__000046_State.Structname = `OpsLine`
	__Field__000046_State.Fieldtypename = `OperationalLineStateEnum`

	// Field values setup
	__Field__000047_TargetHeading.Name = `TargetHeading`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetHeading]
	__Field__000047_TargetHeading.Identifier = `ref_models.Liner.TargetHeading`
	__Field__000047_TargetHeading.FieldTypeAsString = ``
	__Field__000047_TargetHeading.Structname = `Liner`
	__Field__000047_TargetHeading.Fieldtypename = `float64`

	// Field values setup
	__Field__000048_TargetLocationLat.Name = `TargetLocationLat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetLocationLat]
	__Field__000048_TargetLocationLat.Identifier = `ref_models.Liner.TargetLocationLat`
	__Field__000048_TargetLocationLat.FieldTypeAsString = ``
	__Field__000048_TargetLocationLat.Structname = `Liner`
	__Field__000048_TargetLocationLat.Fieldtypename = `float64`

	// Field values setup
	__Field__000049_TargetLocationLat.Name = `TargetLocationLat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.TargetLocationLat]
	__Field__000049_TargetLocationLat.Identifier = `ref_models.Message.TargetLocationLat`
	__Field__000049_TargetLocationLat.FieldTypeAsString = ``
	__Field__000049_TargetLocationLat.Structname = `Message`
	__Field__000049_TargetLocationLat.Fieldtypename = `float64`

	// Field values setup
	__Field__000050_TargetLocationLng.Name = `TargetLocationLng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.TargetLocationLng]
	__Field__000050_TargetLocationLng.Identifier = `ref_models.Message.TargetLocationLng`
	__Field__000050_TargetLocationLng.FieldTypeAsString = ``
	__Field__000050_TargetLocationLng.Structname = `Message`
	__Field__000050_TargetLocationLng.Fieldtypename = `float64`

	// Field values setup
	__Field__000051_TargetLocationLng.Name = `TargetLocationLng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetLocationLng]
	__Field__000051_TargetLocationLng.Identifier = `ref_models.Liner.TargetLocationLng`
	__Field__000051_TargetLocationLng.FieldTypeAsString = ``
	__Field__000051_TargetLocationLng.Structname = `Liner`
	__Field__000051_TargetLocationLng.Fieldtypename = `float64`

	// Field values setup
	__Field__000052_Timestampstartstring.Name = `Timestampstartstring`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Timestampstartstring]
	__Field__000052_Timestampstartstring.Identifier = `ref_models.Message.Timestampstartstring`
	__Field__000052_Timestampstartstring.FieldTypeAsString = ``
	__Field__000052_Timestampstartstring.Structname = `Message`
	__Field__000052_Timestampstartstring.Fieldtypename = `string`

	// Field values setup
	__Field__000053_Timestampstring.Name = `Timestampstring`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message.Timestampstring]
	__Field__000053_Timestampstring.Identifier = `ref_models.Message.Timestampstring`
	__Field__000053_Timestampstring.FieldTypeAsString = ``
	__Field__000053_Timestampstring.Structname = `Message`
	__Field__000053_Timestampstring.Fieldtypename = `string`

	// Field values setup
	__Field__000054_Timestampstring.Name = `Timestampstring`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Timestampstring]
	__Field__000054_Timestampstring.Identifier = `ref_models.Liner.Timestampstring`
	__Field__000054_Timestampstring.FieldTypeAsString = ``
	__Field__000054_Timestampstring.Structname = `Liner`
	__Field__000054_Timestampstring.Fieldtypename = `string`

	// Field values setup
	__Field__000055_Timestampstring.Name = `Timestampstring`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Timestampstring]
	__Field__000055_Timestampstring.Identifier = `ref_models.Satellite.Timestampstring`
	__Field__000055_Timestampstring.FieldTypeAsString = ``
	__Field__000055_Timestampstring.Structname = `Satellite`
	__Field__000055_Timestampstring.Fieldtypename = `string`

	// Field values setup
	__Field__000056_TransmissionMessage.Name = `TransmissionMessage`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.TransmissionMessage]
	__Field__000056_TransmissionMessage.Identifier = `ref_models.OpsLine.TransmissionMessage`
	__Field__000056_TransmissionMessage.FieldTypeAsString = ``
	__Field__000056_TransmissionMessage.Structname = `OpsLine`
	__Field__000056_TransmissionMessage.Fieldtypename = `string`

	// Field values setup
	__Field__000057_TransmissionMessageBackward.Name = `TransmissionMessageBackward`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.TransmissionMessageBackward]
	__Field__000057_TransmissionMessageBackward.Identifier = `ref_models.OpsLine.TransmissionMessageBackward`
	__Field__000057_TransmissionMessageBackward.FieldTypeAsString = ``
	__Field__000057_TransmissionMessageBackward.Structname = `OpsLine`
	__Field__000057_TransmissionMessageBackward.Fieldtypename = `string`

	// Field values setup
	__Field__000058_VerticalSpeed.Name = `VerticalSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.VerticalSpeed]
	__Field__000058_VerticalSpeed.Identifier = `ref_models.Liner.VerticalSpeed`
	__Field__000058_VerticalSpeed.FieldTypeAsString = ``
	__Field__000058_VerticalSpeed.Structname = `Liner`
	__Field__000058_VerticalSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000059_VerticalSpeed.Name = `VerticalSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.VerticalSpeed]
	__Field__000059_VerticalSpeed.Identifier = `ref_models.Satellite.VerticalSpeed`
	__Field__000059_VerticalSpeed.FieldTypeAsString = ``
	__Field__000059_VerticalSpeed.Structname = `Satellite`
	__Field__000059_VerticalSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000060_ZoomLevel.Name = `ZoomLevel`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario.ZoomLevel]
	__Field__000060_ZoomLevel.Identifier = `ref_models.Scenario.ZoomLevel`
	__Field__000060_ZoomLevel.FieldTypeAsString = ``
	__Field__000060_ZoomLevel.Structname = `Scenario`
	__Field__000060_ZoomLevel.Fieldtypename = `float64`

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_ConceptEnum.Name = `NewDiagram-ConceptEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ConceptEnum]
	__GongEnumShape__000000_NewDiagram_ConceptEnum.Identifier = `ref_models.ConceptEnum`
	__GongEnumShape__000000_NewDiagram_ConceptEnum.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_ConceptEnum.Heigth = 63.000000

	// GongEnumShape values setup
	__GongEnumShape__000001_NewDiagram_LinerStateEnum.Name = `NewDiagram-LinerStateEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinerStateEnum]
	__GongEnumShape__000001_NewDiagram_LinerStateEnum.Identifier = `ref_models.LinerStateEnum`
	__GongEnumShape__000001_NewDiagram_LinerStateEnum.Width = 240.000000
	__GongEnumShape__000001_NewDiagram_LinerStateEnum.Heigth = 63.000000

	// GongEnumShape values setup
	__GongEnumShape__000002_NewDiagram_MessageStateEnum.Name = `NewDiagram-MessageStateEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MessageStateEnum]
	__GongEnumShape__000002_NewDiagram_MessageStateEnum.Identifier = `ref_models.MessageStateEnum`
	__GongEnumShape__000002_NewDiagram_MessageStateEnum.Width = 240.000000
	__GongEnumShape__000002_NewDiagram_MessageStateEnum.Heigth = 63.000000

	// GongEnumShape values setup
	__GongEnumShape__000003_NewDiagram_OperationalLineStateEnum.Name = `NewDiagram-OperationalLineStateEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OperationalLineStateEnum]
	__GongEnumShape__000003_NewDiagram_OperationalLineStateEnum.Identifier = `ref_models.OperationalLineStateEnum`
	__GongEnumShape__000003_NewDiagram_OperationalLineStateEnum.Width = 240.000000
	__GongEnumShape__000003_NewDiagram_OperationalLineStateEnum.Heigth = 63.000000

	// GongEnumShape values setup
	__GongEnumShape__000004_NewDiagram_OrderEnum.Name = `NewDiagram-OrderEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OrderEnum]
	__GongEnumShape__000004_NewDiagram_OrderEnum.Identifier = `ref_models.OrderEnum`
	__GongEnumShape__000004_NewDiagram_OrderEnum.Width = 240.000000
	__GongEnumShape__000004_NewDiagram_OrderEnum.Heigth = 63.000000

	// GongEnumShape values setup
	__GongEnumShape__000005_NewDiagram_RadarStateEnum.Name = `NewDiagram-RadarStateEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RadarStateEnum]
	__GongEnumShape__000005_NewDiagram_RadarStateEnum.Identifier = `ref_models.RadarStateEnum`
	__GongEnumShape__000005_NewDiagram_RadarStateEnum.Width = 240.000000
	__GongEnumShape__000005_NewDiagram_RadarStateEnum.Heigth = 63.000000

	// GongEnumShape values setup
	__GongEnumShape__000006_NewDiagram_ReportEnum.Name = `NewDiagram-ReportEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ReportEnum]
	__GongEnumShape__000006_NewDiagram_ReportEnum.Identifier = `ref_models.ReportEnum`
	__GongEnumShape__000006_NewDiagram_ReportEnum.Width = 240.000000
	__GongEnumShape__000006_NewDiagram_ReportEnum.Heigth = 63.000000

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_CivilianAirport.Name = `NewDiagram-CivilianAirport`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport]
	__GongStructShape__000000_NewDiagram_CivilianAirport.Identifier = `ref_models.CivilianAirport`
	__GongStructShape__000000_NewDiagram_CivilianAirport.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_CivilianAirport.NbInstances = 2
	__GongStructShape__000000_NewDiagram_CivilianAirport.Width = 240.000000
	__GongStructShape__000000_NewDiagram_CivilianAirport.Heigth = 108.000000
	__GongStructShape__000000_NewDiagram_CivilianAirport.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Liner.Name = `NewDiagram-Liner`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner]
	__GongStructShape__000001_NewDiagram_Liner.Identifier = `ref_models.Liner`
	__GongStructShape__000001_NewDiagram_Liner.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_Liner.NbInstances = 1
	__GongStructShape__000001_NewDiagram_Liner.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Liner.Heigth = 273.000000
	__GongStructShape__000001_NewDiagram_Liner.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Message.Name = `NewDiagram-Message`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message]
	__GongStructShape__000002_NewDiagram_Message.Identifier = `ref_models.Message`
	__GongStructShape__000002_NewDiagram_Message.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_Message.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Message.Width = 377.000000
	__GongStructShape__000002_NewDiagram_Message.Heigth = 333.000000
	__GongStructShape__000002_NewDiagram_Message.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_OpsLine.Name = `NewDiagram-OpsLine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine]
	__GongStructShape__000003_NewDiagram_OpsLine.Identifier = `ref_models.OpsLine`
	__GongStructShape__000003_NewDiagram_OpsLine.ShowNbInstances = true
	__GongStructShape__000003_NewDiagram_OpsLine.NbInstances = 0
	__GongStructShape__000003_NewDiagram_OpsLine.Width = 329.000000
	__GongStructShape__000003_NewDiagram_OpsLine.Heigth = 245.000000
	__GongStructShape__000003_NewDiagram_OpsLine.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_Radar.Name = `NewDiagram-Radar`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Radar]
	__GongStructShape__000004_NewDiagram_Radar.Identifier = `ref_models.Radar`
	__GongStructShape__000004_NewDiagram_Radar.ShowNbInstances = true
	__GongStructShape__000004_NewDiagram_Radar.NbInstances = 0
	__GongStructShape__000004_NewDiagram_Radar.Width = 240.000000
	__GongStructShape__000004_NewDiagram_Radar.Heigth = 138.000000
	__GongStructShape__000004_NewDiagram_Radar.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NewDiagram_Satellite.Name = `NewDiagram-Satellite`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite]
	__GongStructShape__000005_NewDiagram_Satellite.Identifier = `ref_models.Satellite`
	__GongStructShape__000005_NewDiagram_Satellite.ShowNbInstances = true
	__GongStructShape__000005_NewDiagram_Satellite.NbInstances = 0
	__GongStructShape__000005_NewDiagram_Satellite.Width = 240.000000
	__GongStructShape__000005_NewDiagram_Satellite.Heigth = 213.000000
	__GongStructShape__000005_NewDiagram_Satellite.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NewDiagram_Scenario.Name = `NewDiagram-Scenario`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario]
	__GongStructShape__000006_NewDiagram_Scenario.Identifier = `ref_models.Scenario`
	__GongStructShape__000006_NewDiagram_Scenario.ShowNbInstances = true
	__GongStructShape__000006_NewDiagram_Scenario.NbInstances = 1
	__GongStructShape__000006_NewDiagram_Scenario.Width = 240.000000
	__GongStructShape__000006_NewDiagram_Scenario.Heigth = 138.000000
	__GongStructShape__000006_NewDiagram_Scenario.IsSelected = false

	// Link values setup
	__Link__000000_ReporingLine.Name = `ReporingLine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.ReporingLine]
	__Link__000000_ReporingLine.Identifier = `ref_models.Liner.ReporingLine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine]
	__Link__000000_ReporingLine.Fieldtypename = `ref_models.OpsLine`
	__Link__000000_ReporingLine.FieldOffsetX = -101.999969
	__Link__000000_ReporingLine.FieldOffsetY = -15.000000
	__Link__000000_ReporingLine.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_ReporingLine.TargetMultiplicityOffsetX = -46.999969
	__Link__000000_ReporingLine.TargetMultiplicityOffsetY = 24.000000
	__Link__000000_ReporingLine.SourceMultiplicity = models.MANY
	__Link__000000_ReporingLine.SourceMultiplicityOffsetX = 19.000000
	__Link__000000_ReporingLine.SourceMultiplicityOffsetY = 23.000000
	__Link__000000_ReporingLine.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ReporingLine.StartRatio = 0.520147
	__Link__000000_ReporingLine.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ReporingLine.EndRatio = 0.438710
	__Link__000000_ReporingLine.CornerOffsetRatio = 1.374236

	// Link values setup
	__Link__000001_Scenario.Name = `Scenario`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.Scenario]
	__Link__000001_Scenario.Identifier = `ref_models.OpsLine.Scenario`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario]
	__Link__000001_Scenario.Fieldtypename = `ref_models.Scenario`
	__Link__000001_Scenario.FieldOffsetX = -75.000000
	__Link__000001_Scenario.FieldOffsetY = -19.000000
	__Link__000001_Scenario.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Scenario.TargetMultiplicityOffsetX = -36.000000
	__Link__000001_Scenario.TargetMultiplicityOffsetY = 28.000000
	__Link__000001_Scenario.SourceMultiplicity = models.MANY
	__Link__000001_Scenario.SourceMultiplicityOffsetX = 22.000000
	__Link__000001_Scenario.SourceMultiplicityOffsetY = 37.000000
	__Link__000001_Scenario.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Scenario.StartRatio = 0.500000
	__Link__000001_Scenario.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Scenario.EndRatio = 0.500000
	__Link__000001_Scenario.CornerOffsetRatio = 1.242604

	// Position values setup
	__Position__000000_Pos_NewDiagram_CivilianAirport.X = 33.000000
	__Position__000000_Pos_NewDiagram_CivilianAirport.Y = 81.000000
	__Position__000000_Pos_NewDiagram_CivilianAirport.Name = `Pos-NewDiagram-CivilianAirport`

	// Position values setup
	__Position__000001_Pos_NewDiagram_ConceptEnum.X = 1054.999969
	__Position__000001_Pos_NewDiagram_ConceptEnum.Y = 628.000000
	__Position__000001_Pos_NewDiagram_ConceptEnum.Name = `Pos-NewDiagram-ConceptEnum`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Liner.X = 38.000000
	__Position__000002_Pos_NewDiagram_Liner.Y = 221.000000
	__Position__000002_Pos_NewDiagram_Liner.Name = `Pos-NewDiagram-Liner`

	// Position values setup
	__Position__000003_Pos_NewDiagram_LinerStateEnum.X = 1044.999969
	__Position__000003_Pos_NewDiagram_LinerStateEnum.Y = 199.000000
	__Position__000003_Pos_NewDiagram_LinerStateEnum.Name = `Pos-NewDiagram-LinerStateEnum`

	// Position values setup
	__Position__000004_Pos_NewDiagram_Message.X = 531.000000
	__Position__000004_Pos_NewDiagram_Message.Y = 64.000000
	__Position__000004_Pos_NewDiagram_Message.Name = `Pos-NewDiagram-Message`

	// Position values setup
	__Position__000005_Pos_NewDiagram_MessageStateEnum.X = 1052.999969
	__Position__000005_Pos_NewDiagram_MessageStateEnum.Y = 523.000000
	__Position__000005_Pos_NewDiagram_MessageStateEnum.Name = `Pos-NewDiagram-MessageStateEnum`

	// Position values setup
	__Position__000006_Pos_NewDiagram_OperationalLineStateEnum.X = 1042.999969
	__Position__000006_Pos_NewDiagram_OperationalLineStateEnum.Y = 113.000000
	__Position__000006_Pos_NewDiagram_OperationalLineStateEnum.Name = `Pos-NewDiagram-OperationalLineStateEnum`

	// Position values setup
	__Position__000007_Pos_NewDiagram_OpsLine.X = 533.000000
	__Position__000007_Pos_NewDiagram_OpsLine.Y = 439.000000
	__Position__000007_Pos_NewDiagram_OpsLine.Name = `Pos-NewDiagram-OpsLine`

	// Position values setup
	__Position__000008_Pos_NewDiagram_OrderEnum.X = 1047.999969
	__Position__000008_Pos_NewDiagram_OrderEnum.Y = 412.000000
	__Position__000008_Pos_NewDiagram_OrderEnum.Name = `Pos-NewDiagram-OrderEnum`

	// Position values setup
	__Position__000009_Pos_NewDiagram_Radar.X = 39.000031
	__Position__000009_Pos_NewDiagram_Radar.Y = 523.000000
	__Position__000009_Pos_NewDiagram_Radar.Name = `Pos-NewDiagram-Radar`

	// Position values setup
	__Position__000010_Pos_NewDiagram_RadarStateEnum.X = 1046.999969
	__Position__000010_Pos_NewDiagram_RadarStateEnum.Y = 304.000000
	__Position__000010_Pos_NewDiagram_RadarStateEnum.Name = `Pos-NewDiagram-RadarStateEnum`

	// Position values setup
	__Position__000011_Pos_NewDiagram_ReportEnum.X = 1043.999969
	__Position__000011_Pos_NewDiagram_ReportEnum.Y = 16.000000
	__Position__000011_Pos_NewDiagram_ReportEnum.Name = `Pos-NewDiagram-ReportEnum`

	// Position values setup
	__Position__000012_Pos_NewDiagram_Satellite.X = 50.000031
	__Position__000012_Pos_NewDiagram_Satellite.Y = 733.000000
	__Position__000012_Pos_NewDiagram_Satellite.Name = `Pos-NewDiagram-Satellite`

	// Position values setup
	__Position__000013_Pos_NewDiagram_Scenario.X = 1064.000000
	__Position__000013_Pos_NewDiagram_Scenario.Y = 738.000000
	__Position__000013_Pos_NewDiagram_Scenario.Name = `Pos-NewDiagram-Scenario`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine.X = 400.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine.Y = 604.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Liner and NewDiagram-OpsLine`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_OpsLine_and_NewDiagram_Scenario.X = 886.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_OpsLine_and_NewDiagram_Scenario.Y = 377.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_OpsLine_and_NewDiagram_Scenario.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-OpsLine and NewDiagram-Scenario`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_CivilianAirport)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Liner)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Message)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_OpsLine)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_Radar)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000005_NewDiagram_Satellite)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_Scenario)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000001_NewDiagram_LinerStateEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000000_NewDiagram_ConceptEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000002_NewDiagram_MessageStateEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000004_NewDiagram_OrderEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000005_NewDiagram_RadarStateEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000003_NewDiagram_OperationalLineStateEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000006_NewDiagram_ReportEnum)
	__GongEnumShape__000000_NewDiagram_ConceptEnum.Position = __Position__000001_Pos_NewDiagram_ConceptEnum
	__GongEnumShape__000001_NewDiagram_LinerStateEnum.Position = __Position__000003_Pos_NewDiagram_LinerStateEnum
	__GongEnumShape__000002_NewDiagram_MessageStateEnum.Position = __Position__000005_Pos_NewDiagram_MessageStateEnum
	__GongEnumShape__000003_NewDiagram_OperationalLineStateEnum.Position = __Position__000006_Pos_NewDiagram_OperationalLineStateEnum
	__GongEnumShape__000004_NewDiagram_OrderEnum.Position = __Position__000008_Pos_NewDiagram_OrderEnum
	__GongEnumShape__000005_NewDiagram_RadarStateEnum.Position = __Position__000010_Pos_NewDiagram_RadarStateEnum
	__GongEnumShape__000006_NewDiagram_ReportEnum.Position = __Position__000011_Pos_NewDiagram_ReportEnum
	__GongStructShape__000000_NewDiagram_CivilianAirport.Position = __Position__000000_Pos_NewDiagram_CivilianAirport
	__GongStructShape__000000_NewDiagram_CivilianAirport.Fields = append(__GongStructShape__000000_NewDiagram_CivilianAirport.Fields, __Field__000015_Lat)
	__GongStructShape__000000_NewDiagram_CivilianAirport.Fields = append(__GongStructShape__000000_NewDiagram_CivilianAirport.Fields, __Field__000023_Lng)
	__GongStructShape__000000_NewDiagram_CivilianAirport.Fields = append(__GongStructShape__000000_NewDiagram_CivilianAirport.Fields, __Field__000034_Name)
	__GongStructShape__000001_NewDiagram_Liner.Position = __Position__000002_Pos_NewDiagram_Liner
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000035_Name)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000013_Lat)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000025_Lng)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000007_Heading)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000020_Level)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000041_Speed)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000045_State)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000047_TargetHeading)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000048_TargetLocationLat)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000051_TargetLocationLng)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000004_DistanceToTarget)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000029_MaxRotationalSpeed)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000058_VerticalSpeed)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000054_Timestampstring)
	__GongStructShape__000001_NewDiagram_Liner.Links = append(__GongStructShape__000001_NewDiagram_Liner.Links, __Link__000000_ReporingLine)
	__GongStructShape__000002_NewDiagram_Message.Position = __Position__000004_Pos_NewDiagram_Message
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000016_Lat)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000024_Lng)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000008_Heading)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000019_Level)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000040_Speed)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000044_State)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000031_Name)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000049_TargetLocationLat)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000050_TargetLocationLng)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000005_DistanceToTarget)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000053_Timestampstring)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000006_DurationSinceSimulationStart)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000052_Timestampstartstring)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000039_Source)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000002_Destination)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000001_Content)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000000_About_string)
	__GongStructShape__000002_NewDiagram_Message.Fields = append(__GongStructShape__000002_NewDiagram_Message.Fields, __Field__000003_Display)
	__GongStructShape__000003_NewDiagram_OpsLine.Position = __Position__000007_Pos_NewDiagram_OpsLine
	__GongStructShape__000003_NewDiagram_OpsLine.Fields = append(__GongStructShape__000003_NewDiagram_OpsLine.Fields, __Field__000010_IsTransmitting)
	__GongStructShape__000003_NewDiagram_OpsLine.Fields = append(__GongStructShape__000003_NewDiagram_OpsLine.Fields, __Field__000056_TransmissionMessage)
	__GongStructShape__000003_NewDiagram_OpsLine.Fields = append(__GongStructShape__000003_NewDiagram_OpsLine.Fields, __Field__000011_IsTransmittingBackward)
	__GongStructShape__000003_NewDiagram_OpsLine.Fields = append(__GongStructShape__000003_NewDiagram_OpsLine.Fields, __Field__000057_TransmissionMessageBackward)
	__GongStructShape__000003_NewDiagram_OpsLine.Fields = append(__GongStructShape__000003_NewDiagram_OpsLine.Fields, __Field__000046_State)
	__GongStructShape__000003_NewDiagram_OpsLine.Fields = append(__GongStructShape__000003_NewDiagram_OpsLine.Fields, __Field__000033_Name)
	__GongStructShape__000003_NewDiagram_OpsLine.Links = append(__GongStructShape__000003_NewDiagram_OpsLine.Links, __Link__000001_Scenario)
	__GongStructShape__000004_NewDiagram_Radar.Position = __Position__000009_Pos_NewDiagram_Radar
	__GongStructShape__000004_NewDiagram_Radar.Fields = append(__GongStructShape__000004_NewDiagram_Radar.Fields, __Field__000043_State)
	__GongStructShape__000004_NewDiagram_Radar.Fields = append(__GongStructShape__000004_NewDiagram_Radar.Fields, __Field__000036_Name)
	__GongStructShape__000004_NewDiagram_Radar.Fields = append(__GongStructShape__000004_NewDiagram_Radar.Fields, __Field__000014_Lat)
	__GongStructShape__000004_NewDiagram_Radar.Fields = append(__GongStructShape__000004_NewDiagram_Radar.Fields, __Field__000026_Lng)
	__GongStructShape__000004_NewDiagram_Radar.Fields = append(__GongStructShape__000004_NewDiagram_Radar.Fields, __Field__000038_Range)
	__GongStructShape__000005_NewDiagram_Satellite.Position = __Position__000012_Pos_NewDiagram_Satellite
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000032_Name)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000021_Line1)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000022_Line2)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000012_Lat)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000028_Lng)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000009_Heading)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000018_Level)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000042_Speed)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000059_VerticalSpeed)
	__GongStructShape__000005_NewDiagram_Satellite.Fields = append(__GongStructShape__000005_NewDiagram_Satellite.Fields, __Field__000055_Timestampstring)
	__GongStructShape__000006_NewDiagram_Scenario.Position = __Position__000013_Pos_NewDiagram_Scenario
	__GongStructShape__000006_NewDiagram_Scenario.Fields = append(__GongStructShape__000006_NewDiagram_Scenario.Fields, __Field__000037_Name)
	__GongStructShape__000006_NewDiagram_Scenario.Fields = append(__GongStructShape__000006_NewDiagram_Scenario.Fields, __Field__000017_Lat)
	__GongStructShape__000006_NewDiagram_Scenario.Fields = append(__GongStructShape__000006_NewDiagram_Scenario.Fields, __Field__000027_Lng)
	__GongStructShape__000006_NewDiagram_Scenario.Fields = append(__GongStructShape__000006_NewDiagram_Scenario.Fields, __Field__000060_ZoomLevel)
	__GongStructShape__000006_NewDiagram_Scenario.Fields = append(__GongStructShape__000006_NewDiagram_Scenario.Fields, __Field__000030_MessageVisualSpeed)
	__Link__000000_ReporingLine.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine
	__Link__000001_Scenario.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_OpsLine_and_NewDiagram_Scenario
}


