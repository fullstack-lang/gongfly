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

	"ref_models.CivilianAirport": &(ref_models.CivilianAirport{}),

	"ref_models.CivilianAirport.Lat": (ref_models.CivilianAirport{}).Lat,

	"ref_models.CivilianAirport.Lng": (ref_models.CivilianAirport{}).Lng,

	"ref_models.CivilianAirport.Name": (ref_models.CivilianAirport{}).Name,

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

	"ref_models.Message": &(ref_models.Message{}),

	"ref_models.OpsLine": &(ref_models.OpsLine{}),

	"ref_models.Radar": &(ref_models.Radar{}),

	"ref_models.Satellite": &(ref_models.Satellite{}),

	"ref_models.Scenario": &(ref_models.Scenario{}),
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
	__Field__000000_DistanceToTarget := (&models.Field{Name: `DistanceToTarget`}).Stage(stage)
	__Field__000001_Heading := (&models.Field{Name: `Heading`}).Stage(stage)
	__Field__000002_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000003_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000004_Level := (&models.Field{Name: `Level`}).Stage(stage)
	__Field__000005_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000006_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000007_MaxRotationalSpeed := (&models.Field{Name: `MaxRotationalSpeed`}).Stage(stage)
	__Field__000008_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000009_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000010_Speed := (&models.Field{Name: `Speed`}).Stage(stage)
	__Field__000011_State := (&models.Field{Name: `State`}).Stage(stage)
	__Field__000012_TargetHeading := (&models.Field{Name: `TargetHeading`}).Stage(stage)
	__Field__000013_TargetLocationLat := (&models.Field{Name: `TargetLocationLat`}).Stage(stage)
	__Field__000014_TargetLocationLng := (&models.Field{Name: `TargetLocationLng`}).Stage(stage)
	__Field__000015_Timestampstring := (&models.Field{Name: `Timestampstring`}).Stage(stage)
	__Field__000016_VerticalSpeed := (&models.Field{Name: `VerticalSpeed`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

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

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_CivilianAirport := (&models.Position{Name: `Pos-NewDiagram-CivilianAirport`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Liner := (&models.Position{Name: `Pos-NewDiagram-Liner`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Message := (&models.Position{Name: `Pos-NewDiagram-Message`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_OpsLine := (&models.Position{Name: `Pos-NewDiagram-OpsLine`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_Radar := (&models.Position{Name: `Pos-NewDiagram-Radar`}).Stage(stage)
	__Position__000005_Pos_NewDiagram_Satellite := (&models.Position{Name: `Pos-NewDiagram-Satellite`}).Stage(stage)
	__Position__000006_Pos_NewDiagram_Scenario := (&models.Position{Name: `Pos-NewDiagram-Scenario`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Liner and NewDiagram-OpsLine`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_DistanceToTarget.Name = `DistanceToTarget`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.DistanceToTarget]
	__Field__000000_DistanceToTarget.Identifier = `ref_models.Liner.DistanceToTarget`
	__Field__000000_DistanceToTarget.FieldTypeAsString = ``
	__Field__000000_DistanceToTarget.Structname = `Liner`
	__Field__000000_DistanceToTarget.Fieldtypename = `float64`

	// Field values setup
	__Field__000001_Heading.Name = `Heading`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Heading]
	__Field__000001_Heading.Identifier = `ref_models.Liner.Heading`
	__Field__000001_Heading.FieldTypeAsString = ``
	__Field__000001_Heading.Structname = `Liner`
	__Field__000001_Heading.Fieldtypename = `float64`

	// Field values setup
	__Field__000002_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Lat]
	__Field__000002_Lat.Identifier = `ref_models.Liner.Lat`
	__Field__000002_Lat.FieldTypeAsString = ``
	__Field__000002_Lat.Structname = `Liner`
	__Field__000002_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000003_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport.Lat]
	__Field__000003_Lat.Identifier = `ref_models.CivilianAirport.Lat`
	__Field__000003_Lat.FieldTypeAsString = ``
	__Field__000003_Lat.Structname = `CivilianAirport`
	__Field__000003_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000004_Level.Name = `Level`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Level]
	__Field__000004_Level.Identifier = `ref_models.Liner.Level`
	__Field__000004_Level.FieldTypeAsString = ``
	__Field__000004_Level.Structname = `Liner`
	__Field__000004_Level.Fieldtypename = `float64`

	// Field values setup
	__Field__000005_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Lng]
	__Field__000005_Lng.Identifier = `ref_models.Liner.Lng`
	__Field__000005_Lng.FieldTypeAsString = ``
	__Field__000005_Lng.Structname = `Liner`
	__Field__000005_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000006_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport.Lng]
	__Field__000006_Lng.Identifier = `ref_models.CivilianAirport.Lng`
	__Field__000006_Lng.FieldTypeAsString = ``
	__Field__000006_Lng.Structname = `CivilianAirport`
	__Field__000006_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000007_MaxRotationalSpeed.Name = `MaxRotationalSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.MaxRotationalSpeed]
	__Field__000007_MaxRotationalSpeed.Identifier = `ref_models.Liner.MaxRotationalSpeed`
	__Field__000007_MaxRotationalSpeed.FieldTypeAsString = ``
	__Field__000007_MaxRotationalSpeed.Structname = `Liner`
	__Field__000007_MaxRotationalSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000008_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport.Name]
	__Field__000008_Name.Identifier = `ref_models.CivilianAirport.Name`
	__Field__000008_Name.FieldTypeAsString = ``
	__Field__000008_Name.Structname = `CivilianAirport`
	__Field__000008_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000009_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Name]
	__Field__000009_Name.Identifier = `ref_models.Liner.Name`
	__Field__000009_Name.FieldTypeAsString = ``
	__Field__000009_Name.Structname = `Liner`
	__Field__000009_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000010_Speed.Name = `Speed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Speed]
	__Field__000010_Speed.Identifier = `ref_models.Liner.Speed`
	__Field__000010_Speed.FieldTypeAsString = ``
	__Field__000010_Speed.Structname = `Liner`
	__Field__000010_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000011_State.Name = `State`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.State]
	__Field__000011_State.Identifier = `ref_models.Liner.State`
	__Field__000011_State.FieldTypeAsString = ``
	__Field__000011_State.Structname = `Liner`
	__Field__000011_State.Fieldtypename = `LinerStateEnum`

	// Field values setup
	__Field__000012_TargetHeading.Name = `TargetHeading`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetHeading]
	__Field__000012_TargetHeading.Identifier = `ref_models.Liner.TargetHeading`
	__Field__000012_TargetHeading.FieldTypeAsString = ``
	__Field__000012_TargetHeading.Structname = `Liner`
	__Field__000012_TargetHeading.Fieldtypename = `float64`

	// Field values setup
	__Field__000013_TargetLocationLat.Name = `TargetLocationLat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetLocationLat]
	__Field__000013_TargetLocationLat.Identifier = `ref_models.Liner.TargetLocationLat`
	__Field__000013_TargetLocationLat.FieldTypeAsString = ``
	__Field__000013_TargetLocationLat.Structname = `Liner`
	__Field__000013_TargetLocationLat.Fieldtypename = `float64`

	// Field values setup
	__Field__000014_TargetLocationLng.Name = `TargetLocationLng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetLocationLng]
	__Field__000014_TargetLocationLng.Identifier = `ref_models.Liner.TargetLocationLng`
	__Field__000014_TargetLocationLng.FieldTypeAsString = ``
	__Field__000014_TargetLocationLng.Structname = `Liner`
	__Field__000014_TargetLocationLng.Fieldtypename = `float64`

	// Field values setup
	__Field__000015_Timestampstring.Name = `Timestampstring`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Timestampstring]
	__Field__000015_Timestampstring.Identifier = `ref_models.Liner.Timestampstring`
	__Field__000015_Timestampstring.FieldTypeAsString = ``
	__Field__000015_Timestampstring.Structname = `Liner`
	__Field__000015_Timestampstring.Fieldtypename = `string`

	// Field values setup
	__Field__000016_VerticalSpeed.Name = `VerticalSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.VerticalSpeed]
	__Field__000016_VerticalSpeed.Identifier = `ref_models.Liner.VerticalSpeed`
	__Field__000016_VerticalSpeed.FieldTypeAsString = ``
	__Field__000016_VerticalSpeed.Structname = `Liner`
	__Field__000016_VerticalSpeed.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_CivilianAirport.Name = `NewDiagram-CivilianAirport`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CivilianAirport]
	__GongStructShape__000000_NewDiagram_CivilianAirport.Identifier = `ref_models.CivilianAirport`
	__GongStructShape__000000_NewDiagram_CivilianAirport.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_CivilianAirport.NbInstances = 0
	__GongStructShape__000000_NewDiagram_CivilianAirport.Width = 240.000000
	__GongStructShape__000000_NewDiagram_CivilianAirport.Heigth = 108.000000
	__GongStructShape__000000_NewDiagram_CivilianAirport.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Liner.Name = `NewDiagram-Liner`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner]
	__GongStructShape__000001_NewDiagram_Liner.Identifier = `ref_models.Liner`
	__GongStructShape__000001_NewDiagram_Liner.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Liner.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Liner.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Liner.Heigth = 273.000000
	__GongStructShape__000001_NewDiagram_Liner.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Message.Name = `NewDiagram-Message`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Message]
	__GongStructShape__000002_NewDiagram_Message.Identifier = `ref_models.Message`
	__GongStructShape__000002_NewDiagram_Message.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_Message.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Message.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Message.Heigth = 63.000000
	__GongStructShape__000002_NewDiagram_Message.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_OpsLine.Name = `NewDiagram-OpsLine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine]
	__GongStructShape__000003_NewDiagram_OpsLine.Identifier = `ref_models.OpsLine`
	__GongStructShape__000003_NewDiagram_OpsLine.ShowNbInstances = false
	__GongStructShape__000003_NewDiagram_OpsLine.NbInstances = 0
	__GongStructShape__000003_NewDiagram_OpsLine.Width = 240.000000
	__GongStructShape__000003_NewDiagram_OpsLine.Heigth = 63.000000
	__GongStructShape__000003_NewDiagram_OpsLine.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_Radar.Name = `NewDiagram-Radar`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Radar]
	__GongStructShape__000004_NewDiagram_Radar.Identifier = `ref_models.Radar`
	__GongStructShape__000004_NewDiagram_Radar.ShowNbInstances = false
	__GongStructShape__000004_NewDiagram_Radar.NbInstances = 0
	__GongStructShape__000004_NewDiagram_Radar.Width = 240.000000
	__GongStructShape__000004_NewDiagram_Radar.Heigth = 63.000000
	__GongStructShape__000004_NewDiagram_Radar.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NewDiagram_Satellite.Name = `NewDiagram-Satellite`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite]
	__GongStructShape__000005_NewDiagram_Satellite.Identifier = `ref_models.Satellite`
	__GongStructShape__000005_NewDiagram_Satellite.ShowNbInstances = false
	__GongStructShape__000005_NewDiagram_Satellite.NbInstances = 0
	__GongStructShape__000005_NewDiagram_Satellite.Width = 240.000000
	__GongStructShape__000005_NewDiagram_Satellite.Heigth = 63.000000
	__GongStructShape__000005_NewDiagram_Satellite.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NewDiagram_Scenario.Name = `NewDiagram-Scenario`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario]
	__GongStructShape__000006_NewDiagram_Scenario.Identifier = `ref_models.Scenario`
	__GongStructShape__000006_NewDiagram_Scenario.ShowNbInstances = false
	__GongStructShape__000006_NewDiagram_Scenario.NbInstances = 0
	__GongStructShape__000006_NewDiagram_Scenario.Width = 240.000000
	__GongStructShape__000006_NewDiagram_Scenario.Heigth = 63.000000
	__GongStructShape__000006_NewDiagram_Scenario.IsSelected = false

	// Link values setup
	__Link__000000_ReporingLine.Name = `ReporingLine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.ReporingLine]
	__Link__000000_ReporingLine.Identifier = `ref_models.Liner.ReporingLine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine]
	__Link__000000_ReporingLine.Fieldtypename = `ref_models.OpsLine`
	__Link__000000_ReporingLine.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_ReporingLine.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_NewDiagram_CivilianAirport.X = 85.000000
	__Position__000000_Pos_NewDiagram_CivilianAirport.Y = 21.000000
	__Position__000000_Pos_NewDiagram_CivilianAirport.Name = `Pos-NewDiagram-CivilianAirport`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Liner.X = 30.000000
	__Position__000001_Pos_NewDiagram_Liner.Y = 440.000000
	__Position__000001_Pos_NewDiagram_Liner.Name = `Pos-NewDiagram-Liner`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Message.X = 440.000000
	__Position__000002_Pos_NewDiagram_Message.Y = 20.000000
	__Position__000002_Pos_NewDiagram_Message.Name = `Pos-NewDiagram-Message`

	// Position values setup
	__Position__000003_Pos_NewDiagram_OpsLine.X = 490.000000
	__Position__000003_Pos_NewDiagram_OpsLine.Y = 460.000000
	__Position__000003_Pos_NewDiagram_OpsLine.Name = `Pos-NewDiagram-OpsLine`

	// Position values setup
	__Position__000004_Pos_NewDiagram_Radar.X = 510.000000
	__Position__000004_Pos_NewDiagram_Radar.Y = 330.000000
	__Position__000004_Pos_NewDiagram_Radar.Name = `Pos-NewDiagram-Radar`

	// Position values setup
	__Position__000005_Pos_NewDiagram_Satellite.X = 490.000000
	__Position__000005_Pos_NewDiagram_Satellite.Y = 200.000000
	__Position__000005_Pos_NewDiagram_Satellite.Name = `Pos-NewDiagram-Satellite`

	// Position values setup
	__Position__000006_Pos_NewDiagram_Scenario.X = 520.000000
	__Position__000006_Pos_NewDiagram_Scenario.Y = 100.000000
	__Position__000006_Pos_NewDiagram_Scenario.Name = `Pos-NewDiagram-Scenario`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine.X = 400.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine.Y = 604.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Liner and NewDiagram-OpsLine`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_CivilianAirport)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Liner)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Message)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_OpsLine)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_Radar)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000005_NewDiagram_Satellite)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_Scenario)
	__GongStructShape__000000_NewDiagram_CivilianAirport.Position = __Position__000000_Pos_NewDiagram_CivilianAirport
	__GongStructShape__000000_NewDiagram_CivilianAirport.Fields = append(__GongStructShape__000000_NewDiagram_CivilianAirport.Fields, __Field__000003_Lat)
	__GongStructShape__000000_NewDiagram_CivilianAirport.Fields = append(__GongStructShape__000000_NewDiagram_CivilianAirport.Fields, __Field__000006_Lng)
	__GongStructShape__000000_NewDiagram_CivilianAirport.Fields = append(__GongStructShape__000000_NewDiagram_CivilianAirport.Fields, __Field__000008_Name)
	__GongStructShape__000001_NewDiagram_Liner.Position = __Position__000001_Pos_NewDiagram_Liner
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000009_Name)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000002_Lat)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000005_Lng)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000001_Heading)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000004_Level)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000010_Speed)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000011_State)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000012_TargetHeading)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000013_TargetLocationLat)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000014_TargetLocationLng)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000000_DistanceToTarget)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000007_MaxRotationalSpeed)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000016_VerticalSpeed)
	__GongStructShape__000001_NewDiagram_Liner.Fields = append(__GongStructShape__000001_NewDiagram_Liner.Fields, __Field__000015_Timestampstring)
	__GongStructShape__000001_NewDiagram_Liner.Links = append(__GongStructShape__000001_NewDiagram_Liner.Links, __Link__000000_ReporingLine)
	__GongStructShape__000002_NewDiagram_Message.Position = __Position__000002_Pos_NewDiagram_Message
	__GongStructShape__000003_NewDiagram_OpsLine.Position = __Position__000003_Pos_NewDiagram_OpsLine
	__GongStructShape__000004_NewDiagram_Radar.Position = __Position__000004_Pos_NewDiagram_Radar
	__GongStructShape__000005_NewDiagram_Satellite.Position = __Position__000005_Pos_NewDiagram_Satellite
	__GongStructShape__000006_NewDiagram_Scenario.Position = __Position__000006_Pos_NewDiagram_Scenario
	__Link__000000_ReporingLine.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Liner_and_NewDiagram_OpsLine
}
