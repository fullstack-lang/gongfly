package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongfly/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_gongfly models.StageStruct
var ___dummy__Time_gongfly time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_gongfly ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_gongfly map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["gongfly"] = gongflyInjection
// }

// gongflyInjection will stage objects of database "gongfly"
func gongflyInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_gongfly := (&models.Classdiagram{Name: `gongfly`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0000 := (&models.Classshape{Name: `Classshape0000`}).Stage()
	__Classshape__000001_Classshape0001 := (&models.Classshape{Name: `Classshape0001`}).Stage()
	__Classshape__000002_Classshape0002 := (&models.Classshape{Name: `Classshape0002`}).Stage()
	__Classshape__000003_Classshape0003 := (&models.Classshape{Name: `Classshape0003`}).Stage()

	// Declarations of staged instances of DiagramPackage
	__DiagramPackage__000000_gongfly_diagrams := (&models.DiagramPackage{Name: `gongfly_diagrams`}).Stage()

	// Declarations of staged instances of Field
	__Field__000000_DistanceToTarget := (&models.Field{Name: `DistanceToTarget`}).Stage()
	__Field__000001_Heading := (&models.Field{Name: `Heading`}).Stage()
	__Field__000002_IsTransmitting := (&models.Field{Name: `IsTransmitting`}).Stage()
	__Field__000003_IsTransmittingBackward := (&models.Field{Name: `IsTransmittingBackward`}).Stage()
	__Field__000004_Lat := (&models.Field{Name: `Lat`}).Stage()
	__Field__000005_Level := (&models.Field{Name: `Level`}).Stage()
	__Field__000006_Line1 := (&models.Field{Name: `Line1`}).Stage()
	__Field__000007_Line2 := (&models.Field{Name: `Line2`}).Stage()
	__Field__000008_Lng := (&models.Field{Name: `Lng`}).Stage()
	__Field__000009_MaxRotationalSpeed := (&models.Field{Name: `MaxRotationalSpeed`}).Stage()
	__Field__000010_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000011_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000014_Speed := (&models.Field{Name: `Speed`}).Stage()
	__Field__000015_State := (&models.Field{Name: `State`}).Stage()
	__Field__000016_State := (&models.Field{Name: `State`}).Stage()
	__Field__000017_TargetHeading := (&models.Field{Name: `TargetHeading`}).Stage()
	__Field__000018_TargetLocationLat := (&models.Field{Name: `TargetLocationLat`}).Stage()
	__Field__000019_TargetLocationLng := (&models.Field{Name: `TargetLocationLng`}).Stage()
	__Field__000020_Timestampstring := (&models.Field{Name: `Timestampstring`}).Stage()
	__Field__000021_TransmissionMessage := (&models.Field{Name: `TransmissionMessage`}).Stage()
	__Field__000022_TransmissionMessageBackward := (&models.Field{Name: `TransmissionMessageBackward`}).Stage()
	__Field__000023_VerticalSpeed := (&models.Field{Name: `VerticalSpeed`}).Stage()
	__Field__000024_ZoomLevel := (&models.Field{Name: `ZoomLevel`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_ReporingLine := (&models.Link{Name: `ReporingLine`}).Stage()
	__Link__000001_Scenario := (&models.Link{Name: `Scenario`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Position_0000 := (&models.Position{Name: `Position-0000`}).Stage()
	__Position__000001_Position_0001 := (&models.Position{Name: `Position-0001`}).Stage()
	__Position__000002_Position_0002 := (&models.Position{Name: `Position-0002`}).Stage()
	__Position__000003_Position_0003 := (&models.Position{Name: `Position-0003`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Liner := (&models.Reference{Name: `Liner`}).Stage()
	__Reference__000001_OpsLine := (&models.Reference{Name: `OpsLine`}).Stage()
	__Reference__000002_Satellite := (&models.Reference{Name: `Satellite`}).Stage()
	__Reference__000003_Scenario := (&models.Reference{Name: `Scenario`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0000 := (&models.Vertice{Name: `Vertice-0000`}).Stage()
	__Vertice__000001_Vertice_0001 := (&models.Vertice{Name: `Vertice-0001`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_gongfly.Name = `gongfly`
	__Classdiagram__000000_gongfly.IsInDrawMode = false

	// Classshape values setup
	__Classshape__000000_Classshape0000.Name = `Classshape0000`
	__Classshape__000000_Classshape0000.ReferenceName = `Liner`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner]
	__Classshape__000000_Classshape0000.Identifier = `ref_models.Liner`
	__Classshape__000000_Classshape0000.ShowNbInstances = true
	__Classshape__000000_Classshape0000.NbInstances = 0
	__Classshape__000000_Classshape0000.Width = 240.000000
	__Classshape__000000_Classshape0000.Heigth = 273.000000
	__Classshape__000000_Classshape0000.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0001.Name = `Classshape0001`
	__Classshape__000001_Classshape0001.ReferenceName = `OpsLine`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine]
	__Classshape__000001_Classshape0001.Identifier = `ref_models.OpsLine`
	__Classshape__000001_Classshape0001.ShowNbInstances = true
	__Classshape__000001_Classshape0001.NbInstances = 0
	__Classshape__000001_Classshape0001.Width = 240.000000
	__Classshape__000001_Classshape0001.Heigth = 153.000000
	__Classshape__000001_Classshape0001.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0002.Name = `Classshape0002`
	__Classshape__000002_Classshape0002.ReferenceName = `Satellite`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite]
	__Classshape__000002_Classshape0002.Identifier = `ref_models.Satellite`
	__Classshape__000002_Classshape0002.ShowNbInstances = true
	__Classshape__000002_Classshape0002.NbInstances = 0
	__Classshape__000002_Classshape0002.Width = 240.000000
	__Classshape__000002_Classshape0002.Heigth = 108.000000
	__Classshape__000002_Classshape0002.IsSelected = false

	// Classshape values setup
	__Classshape__000003_Classshape0003.Name = `Classshape0003`
	__Classshape__000003_Classshape0003.ReferenceName = `Scenario`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario]
	__Classshape__000003_Classshape0003.Identifier = `ref_models.Scenario`
	__Classshape__000003_Classshape0003.ShowNbInstances = true
	__Classshape__000003_Classshape0003.NbInstances = 0
	__Classshape__000003_Classshape0003.Width = 240.000000
	__Classshape__000003_Classshape0003.Heigth = 78.000000
	__Classshape__000003_Classshape0003.IsSelected = false

	// DiagramPackage values setup
	__DiagramPackage__000000_gongfly_diagrams.Name = `gongfly_diagrams`
	__DiagramPackage__000000_gongfly_diagrams.Path = `github.com/fullstack-lang/gongfly/go/models`
	__DiagramPackage__000000_gongfly_diagrams.GongModelPath = `github.com/fullstack-lang/gongfly/go/models`
	__DiagramPackage__000000_gongfly_diagrams.IsEditable = true
	__DiagramPackage__000000_gongfly_diagrams.IsReloaded = false
	__DiagramPackage__000000_gongfly_diagrams.AbsolutePathToDiagramPackage = `D:\MOSS-Users\peugeot\go\src\github.com\fullstack-lang\gongfly\go\diagrams`

	// Field values setup
	__Field__000000_DistanceToTarget.Name = `DistanceToTarget`
	__Field__000000_DistanceToTarget.Fieldname = `DistanceToTarget`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.DistanceToTarget]
	__Field__000000_DistanceToTarget.Identifier = `ref_models.Liner.DistanceToTarget`
	__Field__000000_DistanceToTarget.FieldTypeAsString = ``
	__Field__000000_DistanceToTarget.Structname = `Liner`
	__Field__000000_DistanceToTarget.Fieldtypename = `float64`

	// Field values setup
	__Field__000001_Heading.Name = `Heading`
	__Field__000001_Heading.Fieldname = `Heading`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Heading]
	__Field__000001_Heading.Identifier = `ref_models.Liner.Heading`
	__Field__000001_Heading.FieldTypeAsString = ``
	__Field__000001_Heading.Structname = `Liner`
	__Field__000001_Heading.Fieldtypename = `float64`

	// Field values setup
	__Field__000002_IsTransmitting.Name = `IsTransmitting`
	__Field__000002_IsTransmitting.Fieldname = `IsTransmitting`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.IsTransmitting]
	__Field__000002_IsTransmitting.Identifier = `ref_models.OpsLine.IsTransmitting`
	__Field__000002_IsTransmitting.FieldTypeAsString = ``
	__Field__000002_IsTransmitting.Structname = `OpsLine`
	__Field__000002_IsTransmitting.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_IsTransmittingBackward.Name = `IsTransmittingBackward`
	__Field__000003_IsTransmittingBackward.Fieldname = `IsTransmittingBackward`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.IsTransmittingBackward]
	__Field__000003_IsTransmittingBackward.Identifier = `ref_models.OpsLine.IsTransmittingBackward`
	__Field__000003_IsTransmittingBackward.FieldTypeAsString = ``
	__Field__000003_IsTransmittingBackward.Structname = `OpsLine`
	__Field__000003_IsTransmittingBackward.Fieldtypename = `bool`

	// Field values setup
	__Field__000004_Lat.Name = `Lat`
	__Field__000004_Lat.Fieldname = `Lat`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Lat]
	__Field__000004_Lat.Identifier = `ref_models.Liner.Lat`
	__Field__000004_Lat.FieldTypeAsString = ``
	__Field__000004_Lat.Structname = `Liner`
	__Field__000004_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000005_Level.Name = `Level`
	__Field__000005_Level.Fieldname = `Level`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Level]
	__Field__000005_Level.Identifier = `ref_models.Liner.Level`
	__Field__000005_Level.FieldTypeAsString = ``
	__Field__000005_Level.Structname = `Liner`
	__Field__000005_Level.Fieldtypename = `float64`

	// Field values setup
	__Field__000006_Line1.Name = `Line1`
	__Field__000006_Line1.Fieldname = `Line1`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Line1]
	__Field__000006_Line1.Identifier = `ref_models.Satellite.Line1`
	__Field__000006_Line1.FieldTypeAsString = ``
	__Field__000006_Line1.Structname = `Satellite`
	__Field__000006_Line1.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Line2.Name = `Line2`
	__Field__000007_Line2.Fieldname = `Line2`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Line2]
	__Field__000007_Line2.Identifier = `ref_models.Satellite.Line2`
	__Field__000007_Line2.FieldTypeAsString = ``
	__Field__000007_Line2.Structname = `Satellite`
	__Field__000007_Line2.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Lng.Name = `Lng`
	__Field__000008_Lng.Fieldname = `Lng`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Lng]
	__Field__000008_Lng.Identifier = `ref_models.Liner.Lng`
	__Field__000008_Lng.FieldTypeAsString = ``
	__Field__000008_Lng.Structname = `Liner`
	__Field__000008_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000009_MaxRotationalSpeed.Name = `MaxRotationalSpeed`
	__Field__000009_MaxRotationalSpeed.Fieldname = `MaxRotationalSpeed`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.MaxRotationalSpeed]
	__Field__000009_MaxRotationalSpeed.Identifier = `ref_models.Liner.MaxRotationalSpeed`
	__Field__000009_MaxRotationalSpeed.FieldTypeAsString = ``
	__Field__000009_MaxRotationalSpeed.Structname = `Liner`
	__Field__000009_MaxRotationalSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000010_Name.Name = `Name`
	__Field__000010_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Name]
	__Field__000010_Name.Identifier = `ref_models.Liner.Name`
	__Field__000010_Name.FieldTypeAsString = ``
	__Field__000010_Name.Structname = `Liner`
	__Field__000010_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000011_Name.Name = `Name`
	__Field__000011_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.Name]
	__Field__000011_Name.Identifier = `ref_models.OpsLine.Name`
	__Field__000011_Name.FieldTypeAsString = ``
	__Field__000011_Name.Structname = `OpsLine`
	__Field__000011_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000012_Name.Name = `Name`
	__Field__000012_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Satellite.Name]
	__Field__000012_Name.Identifier = `ref_models.Satellite.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `Satellite`
	__Field__000012_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000013_Name.Name = `Name`
	__Field__000013_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario.Name]
	__Field__000013_Name.Identifier = `ref_models.Scenario.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Scenario`
	__Field__000013_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000014_Speed.Name = `Speed`
	__Field__000014_Speed.Fieldname = `Speed`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Speed]
	__Field__000014_Speed.Identifier = `ref_models.Liner.Speed`
	__Field__000014_Speed.FieldTypeAsString = ``
	__Field__000014_Speed.Structname = `Liner`
	__Field__000014_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000015_State.Name = `State`
	__Field__000015_State.Fieldname = `State`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.State]
	__Field__000015_State.Identifier = `ref_models.OpsLine.State`
	__Field__000015_State.FieldTypeAsString = ``
	__Field__000015_State.Structname = `OpsLine`
	__Field__000015_State.Fieldtypename = `OperationalLineStateEnum`

	// Field values setup
	__Field__000016_State.Name = `State`
	__Field__000016_State.Fieldname = `State`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.State]
	__Field__000016_State.Identifier = `ref_models.Liner.State`
	__Field__000016_State.FieldTypeAsString = ``
	__Field__000016_State.Structname = `Liner`
	__Field__000016_State.Fieldtypename = `LinerStateEnum`

	// Field values setup
	__Field__000017_TargetHeading.Name = `TargetHeading`
	__Field__000017_TargetHeading.Fieldname = `TargetHeading`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetHeading]
	__Field__000017_TargetHeading.Identifier = `ref_models.Liner.TargetHeading`
	__Field__000017_TargetHeading.FieldTypeAsString = ``
	__Field__000017_TargetHeading.Structname = `Liner`
	__Field__000017_TargetHeading.Fieldtypename = `float64`

	// Field values setup
	__Field__000018_TargetLocationLat.Name = `TargetLocationLat`
	__Field__000018_TargetLocationLat.Fieldname = `TargetLocationLat`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetLocationLat]
	__Field__000018_TargetLocationLat.Identifier = `ref_models.Liner.TargetLocationLat`
	__Field__000018_TargetLocationLat.FieldTypeAsString = ``
	__Field__000018_TargetLocationLat.Structname = `Liner`
	__Field__000018_TargetLocationLat.Fieldtypename = `float64`

	// Field values setup
	__Field__000019_TargetLocationLng.Name = `TargetLocationLng`
	__Field__000019_TargetLocationLng.Fieldname = `TargetLocationLng`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.TargetLocationLng]
	__Field__000019_TargetLocationLng.Identifier = `ref_models.Liner.TargetLocationLng`
	__Field__000019_TargetLocationLng.FieldTypeAsString = ``
	__Field__000019_TargetLocationLng.Structname = `Liner`
	__Field__000019_TargetLocationLng.Fieldtypename = `float64`

	// Field values setup
	__Field__000020_Timestampstring.Name = `Timestampstring`
	__Field__000020_Timestampstring.Fieldname = `Timestampstring`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.Timestampstring]
	__Field__000020_Timestampstring.Identifier = `ref_models.Liner.Timestampstring`
	__Field__000020_Timestampstring.FieldTypeAsString = ``
	__Field__000020_Timestampstring.Structname = `Liner`
	__Field__000020_Timestampstring.Fieldtypename = `string`

	// Field values setup
	__Field__000021_TransmissionMessage.Name = `TransmissionMessage`
	__Field__000021_TransmissionMessage.Fieldname = `TransmissionMessage`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.TransmissionMessage]
	__Field__000021_TransmissionMessage.Identifier = `ref_models.OpsLine.TransmissionMessage`
	__Field__000021_TransmissionMessage.FieldTypeAsString = ``
	__Field__000021_TransmissionMessage.Structname = `OpsLine`
	__Field__000021_TransmissionMessage.Fieldtypename = `string`

	// Field values setup
	__Field__000022_TransmissionMessageBackward.Name = `TransmissionMessageBackward`
	__Field__000022_TransmissionMessageBackward.Fieldname = `TransmissionMessageBackward`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.TransmissionMessageBackward]
	__Field__000022_TransmissionMessageBackward.Identifier = `ref_models.OpsLine.TransmissionMessageBackward`
	__Field__000022_TransmissionMessageBackward.FieldTypeAsString = ``
	__Field__000022_TransmissionMessageBackward.Structname = `OpsLine`
	__Field__000022_TransmissionMessageBackward.Fieldtypename = `string`

	// Field values setup
	__Field__000023_VerticalSpeed.Name = `VerticalSpeed`
	__Field__000023_VerticalSpeed.Fieldname = `VerticalSpeed`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.VerticalSpeed]
	__Field__000023_VerticalSpeed.Identifier = `ref_models.Liner.VerticalSpeed`
	__Field__000023_VerticalSpeed.FieldTypeAsString = ``
	__Field__000023_VerticalSpeed.Structname = `Liner`
	__Field__000023_VerticalSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000024_ZoomLevel.Name = `ZoomLevel`
	__Field__000024_ZoomLevel.Fieldname = `ZoomLevel`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Scenario.ZoomLevel]
	__Field__000024_ZoomLevel.Identifier = `ref_models.Scenario.ZoomLevel`
	__Field__000024_ZoomLevel.FieldTypeAsString = ``
	__Field__000024_ZoomLevel.Structname = `Scenario`
	__Field__000024_ZoomLevel.Fieldtypename = `float64`

	// Link values setup
	__Link__000000_ReporingLine.Name = `ReporingLine`
	__Link__000000_ReporingLine.Fieldname = `ReporingLine`
	__Link__000000_ReporingLine.Structname = `Liner`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Liner.ReporingLine]
	__Link__000000_ReporingLine.Identifier = `ref_models.Liner.ReporingLine`
	__Link__000000_ReporingLine.Fieldtypename = `OpsLine`
	__Link__000000_ReporingLine.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_ReporingLine.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000001_Scenario.Name = `Scenario`
	__Link__000001_Scenario.Fieldname = `Scenario`
	__Link__000001_Scenario.Structname = `OpsLine`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.OpsLine.Scenario]
	__Link__000001_Scenario.Identifier = `ref_models.OpsLine.Scenario`
	__Link__000001_Scenario.Fieldtypename = `Scenario`
	__Link__000001_Scenario.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Scenario.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Position_0000.X = 20.000000
	__Position__000000_Position_0000.Y = 20.000000
	__Position__000000_Position_0000.Name = `Position-0000`

	// Position values setup
	__Position__000001_Position_0001.X = 440.000000
	__Position__000001_Position_0001.Y = 40.000000
	__Position__000001_Position_0001.Name = `Position-0001`

	// Position values setup
	__Position__000002_Position_0002.X = 20.000000
	__Position__000002_Position_0002.Y = 310.000000
	__Position__000002_Position_0002.Name = `Position-0002`

	// Position values setup
	__Position__000003_Position_0003.X = 20.000000
	__Position__000003_Position_0003.Y = 440.000000
	__Position__000003_Position_0003.Name = `Position-0003`

	// Reference values setup
	__Reference__000000_Liner.Name = `Liner`
	__Reference__000000_Liner.NbInstances = 0
	__Reference__000000_Liner.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_OpsLine.Name = `OpsLine`
	__Reference__000001_OpsLine.NbInstances = 0
	__Reference__000001_OpsLine.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000002_Satellite.Name = `Satellite`
	__Reference__000002_Satellite.NbInstances = 0
	__Reference__000002_Satellite.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000003_Scenario.Name = `Scenario`
	__Reference__000003_Scenario.NbInstances = 0
	__Reference__000003_Scenario.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0000.X = 410.000000
	__Vertice__000000_Vertice_0000.Y = 266.500000
	__Vertice__000000_Vertice_0000.Name = `Vertice-0000`

	// Vertice values setup
	__Vertice__000001_Vertice_0001.X = 565.000000
	__Vertice__000001_Vertice_0001.Y = 479.000000
	__Vertice__000001_Vertice_0001.Name = `Vertice-0001`

	// Setup of pointers
	__Classdiagram__000000_gongfly.Classshapes = append(__Classdiagram__000000_gongfly.Classshapes, __Classshape__000000_Classshape0000)
	__Classdiagram__000000_gongfly.Classshapes = append(__Classdiagram__000000_gongfly.Classshapes, __Classshape__000001_Classshape0001)
	__Classdiagram__000000_gongfly.Classshapes = append(__Classdiagram__000000_gongfly.Classshapes, __Classshape__000002_Classshape0002)
	__Classdiagram__000000_gongfly.Classshapes = append(__Classdiagram__000000_gongfly.Classshapes, __Classshape__000003_Classshape0003)
	__Classshape__000000_Classshape0000.Position = __Position__000000_Position_0000
	__Classshape__000000_Classshape0000.Reference = __Reference__000000_Liner
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000000_DistanceToTarget)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000001_Heading)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000004_Lat)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000005_Level)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000008_Lng)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000009_MaxRotationalSpeed)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000010_Name)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000014_Speed)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000016_State)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000017_TargetHeading)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000018_TargetLocationLat)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000019_TargetLocationLng)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000020_Timestampstring)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000023_VerticalSpeed)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000000_ReporingLine)
	__Classshape__000001_Classshape0001.Position = __Position__000001_Position_0001
	__Classshape__000001_Classshape0001.Reference = __Reference__000001_OpsLine
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000002_IsTransmitting)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000003_IsTransmittingBackward)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000011_Name)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000015_State)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000021_TransmissionMessage)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000022_TransmissionMessageBackward)
	__Classshape__000001_Classshape0001.Links = append(__Classshape__000001_Classshape0001.Links, __Link__000001_Scenario)
	__Classshape__000002_Classshape0002.Position = __Position__000002_Position_0002
	__Classshape__000002_Classshape0002.Reference = __Reference__000002_Satellite
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000006_Line1)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000007_Line2)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000012_Name)
	__Classshape__000003_Classshape0003.Position = __Position__000003_Position_0003
	__Classshape__000003_Classshape0003.Reference = __Reference__000003_Scenario
	__Classshape__000003_Classshape0003.Fields = append(__Classshape__000003_Classshape0003.Fields, __Field__000013_Name)
	__Classshape__000003_Classshape0003.Fields = append(__Classshape__000003_Classshape0003.Fields, __Field__000024_ZoomLevel)
	__DiagramPackage__000000_gongfly_diagrams.Classdiagrams = append(__DiagramPackage__000000_gongfly_diagrams.Classdiagrams, __Classdiagram__000000_gongfly)
	__Link__000000_ReporingLine.Middlevertice = __Vertice__000000_Vertice_0000
	__Link__000001_Scenario.Middlevertice = __Vertice__000001_Vertice_0001
}


