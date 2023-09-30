package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongleaflet/go/models"
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

	"ref_models.BACKWARD_END_TO_START": ref_models.BACKWARD_END_TO_START,

	"ref_models.BLUE": ref_models.BLUE,

	"ref_models.Circle": &(ref_models.Circle{}),

	"ref_models.Circle.ColorEnum": (ref_models.Circle{}).ColorEnum,

	"ref_models.Circle.DashStyleEnum": (ref_models.Circle{}).DashStyleEnum,

	"ref_models.Circle.Lat": (ref_models.Circle{}).Lat,

	"ref_models.Circle.LayerGroup": (ref_models.Circle{}).LayerGroup,

	"ref_models.Circle.Lng": (ref_models.Circle{}).Lng,

	"ref_models.Circle.Name": (ref_models.Circle{}).Name,

	"ref_models.Circle.Radius": (ref_models.Circle{}).Radius,

	"ref_models.ColorEnum": ref_models.ColorEnum(""),

	"ref_models.DashStyleEnum": ref_models.DashStyleEnum(""),

	"ref_models.DivIcon": &(ref_models.DivIcon{}),

	"ref_models.DivIcon.Name": (ref_models.DivIcon{}).Name,

	"ref_models.DivIcon.SVG": (ref_models.DivIcon{}).SVG,

	"ref_models.FIVE_TEN": ref_models.FIVE_TEN,

	"ref_models.FIVE_TWENTY": ref_models.FIVE_TWENTY,

	"ref_models.FORWARD_START_TO_END": ref_models.FORWARD_START_TO_END,

	"ref_models.GREEN": ref_models.GREEN,

	"ref_models.GREY": ref_models.GREY,

	"ref_models.IS_NOT_TRANSMITTING": ref_models.IS_NOT_TRANSMITTING,

	"ref_models.IS_TRANSMITTING": ref_models.IS_TRANSMITTING,

	"ref_models.LIGHT_BROWN_8D6E63": ref_models.LIGHT_BROWN_8D6E63,

	"ref_models.LayerGroup": &(ref_models.LayerGroup{}),

	"ref_models.LayerGroup.DisplayName": (ref_models.LayerGroup{}).DisplayName,

	"ref_models.LayerGroup.Name": (ref_models.LayerGroup{}).Name,

	"ref_models.LayerGroupUse": &(ref_models.LayerGroupUse{}),

	"ref_models.LayerGroupUse.Display": (ref_models.LayerGroupUse{}).Display,

	"ref_models.LayerGroupUse.LayerGroup": (ref_models.LayerGroupUse{}).LayerGroup,

	"ref_models.LayerGroupUse.Name": (ref_models.LayerGroupUse{}).Name,

	"ref_models.MapOptions": &(ref_models.MapOptions{}),

	"ref_models.MapOptions.Attribution": (ref_models.MapOptions{}).Attribution,

	"ref_models.MapOptions.AttributionControl": (ref_models.MapOptions{}).AttributionControl,

	"ref_models.MapOptions.Lat": (ref_models.MapOptions{}).Lat,

	"ref_models.MapOptions.LayerGroupUses": (ref_models.MapOptions{}).LayerGroupUses,

	"ref_models.MapOptions.Lng": (ref_models.MapOptions{}).Lng,

	"ref_models.MapOptions.MaxZoom": (ref_models.MapOptions{}).MaxZoom,

	"ref_models.MapOptions.Name": (ref_models.MapOptions{}).Name,

	"ref_models.MapOptions.UrlTemplate": (ref_models.MapOptions{}).UrlTemplate,

	"ref_models.MapOptions.ZoomControl": (ref_models.MapOptions{}).ZoomControl,

	"ref_models.MapOptions.ZoomLevel": (ref_models.MapOptions{}).ZoomLevel,

	"ref_models.MapOptions.ZoomSnap": (ref_models.MapOptions{}).ZoomSnap,

	"ref_models.Marker": &(ref_models.Marker{}),

	"ref_models.Marker.ColorEnum": (ref_models.Marker{}).ColorEnum,

	"ref_models.Marker.DivIcon": (ref_models.Marker{}).DivIcon,

	"ref_models.Marker.Lat": (ref_models.Marker{}).Lat,

	"ref_models.Marker.LayerGroup": (ref_models.Marker{}).LayerGroup,

	"ref_models.Marker.Lng": (ref_models.Marker{}).Lng,

	"ref_models.Marker.Name": (ref_models.Marker{}).Name,

	"ref_models.NONE": ref_models.NONE,

	"ref_models.RED": ref_models.RED,

	"ref_models.Start_To_End_Enum": ref_models.Start_To_End_Enum(""),

	"ref_models.TransmittingEnum": ref_models.TransmittingEnum(""),

	"ref_models.UserClick": &(ref_models.UserClick{}),

	"ref_models.UserClick.Lat": (ref_models.UserClick{}).Lat,

	"ref_models.UserClick.Lng": (ref_models.UserClick{}).Lng,

	"ref_models.UserClick.Name": (ref_models.UserClick{}).Name,

	"ref_models.UserClick.TimeOfClick": (ref_models.UserClick{}).TimeOfClick,

	"ref_models.VLine": &(ref_models.VLine{}),

	"ref_models.VLine.ColorEnum": (ref_models.VLine{}).ColorEnum,

	"ref_models.VLine.DashStyleEnum": (ref_models.VLine{}).DashStyleEnum,

	"ref_models.VLine.EndLat": (ref_models.VLine{}).EndLat,

	"ref_models.VLine.EndLng": (ref_models.VLine{}).EndLng,

	"ref_models.VLine.IsTransmitting": (ref_models.VLine{}).IsTransmitting,

	"ref_models.VLine.IsTransmittingBackward": (ref_models.VLine{}).IsTransmittingBackward,

	"ref_models.VLine.LayerGroup": (ref_models.VLine{}).LayerGroup,

	"ref_models.VLine.Message": (ref_models.VLine{}).Message,

	"ref_models.VLine.MessageBackward": (ref_models.VLine{}).MessageBackward,

	"ref_models.VLine.Name": (ref_models.VLine{}).Name,

	"ref_models.VLine.StartLat": (ref_models.VLine{}).StartLat,

	"ref_models.VLine.StartLng": (ref_models.VLine{}).StartLng,

	"ref_models.VisualTrack": &(ref_models.VisualTrack{}),

	"ref_models.VisualTrack.ColorEnum": (ref_models.VisualTrack{}).ColorEnum,

	"ref_models.VisualTrack.DisplayLevelAndSpeed": (ref_models.VisualTrack{}).DisplayLevelAndSpeed,

	"ref_models.VisualTrack.DisplayTrackHistory": (ref_models.VisualTrack{}).DisplayTrackHistory,

	"ref_models.VisualTrack.DivIcon": (ref_models.VisualTrack{}).DivIcon,

	"ref_models.VisualTrack.Heading": (ref_models.VisualTrack{}).Heading,

	"ref_models.VisualTrack.Lat": (ref_models.VisualTrack{}).Lat,

	"ref_models.VisualTrack.LayerGroup": (ref_models.VisualTrack{}).LayerGroup,

	"ref_models.VisualTrack.Level": (ref_models.VisualTrack{}).Level,

	"ref_models.VisualTrack.Lng": (ref_models.VisualTrack{}).Lng,

	"ref_models.VisualTrack.Name": (ref_models.VisualTrack{}).Name,

	"ref_models.VisualTrack.Speed": (ref_models.VisualTrack{}).Speed,

	"ref_models.VisualTrack.VerticalSpeed": (ref_models.VisualTrack{}).VerticalSpeed,
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
	__Field__000000_Attribution := (&models.Field{Name: `Attribution`}).Stage(stage)
	__Field__000001_AttributionControl := (&models.Field{Name: `AttributionControl`}).Stage(stage)
	__Field__000002_ColorEnum := (&models.Field{Name: `ColorEnum`}).Stage(stage)
	__Field__000003_ColorEnum := (&models.Field{Name: `ColorEnum`}).Stage(stage)
	__Field__000004_ColorEnum := (&models.Field{Name: `ColorEnum`}).Stage(stage)
	__Field__000005_ColorEnum := (&models.Field{Name: `ColorEnum`}).Stage(stage)
	__Field__000006_DashStyleEnum := (&models.Field{Name: `DashStyleEnum`}).Stage(stage)
	__Field__000007_DashStyleEnum := (&models.Field{Name: `DashStyleEnum`}).Stage(stage)
	__Field__000008_Display := (&models.Field{Name: `Display`}).Stage(stage)
	__Field__000009_DisplayLevelAndSpeed := (&models.Field{Name: `DisplayLevelAndSpeed`}).Stage(stage)
	__Field__000010_DisplayName := (&models.Field{Name: `DisplayName`}).Stage(stage)
	__Field__000011_DisplayTrackHistory := (&models.Field{Name: `DisplayTrackHistory`}).Stage(stage)
	__Field__000012_EndLat := (&models.Field{Name: `EndLat`}).Stage(stage)
	__Field__000013_EndLng := (&models.Field{Name: `EndLng`}).Stage(stage)
	__Field__000014_Heading := (&models.Field{Name: `Heading`}).Stage(stage)
	__Field__000015_IsTransmitting := (&models.Field{Name: `IsTransmitting`}).Stage(stage)
	__Field__000016_IsTransmittingBackward := (&models.Field{Name: `IsTransmittingBackward`}).Stage(stage)
	__Field__000017_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000018_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000019_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000020_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000021_Lat := (&models.Field{Name: `Lat`}).Stage(stage)
	__Field__000022_Level := (&models.Field{Name: `Level`}).Stage(stage)
	__Field__000023_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000024_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000025_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000026_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000027_Lng := (&models.Field{Name: `Lng`}).Stage(stage)
	__Field__000028_MaxZoom := (&models.Field{Name: `MaxZoom`}).Stage(stage)
	__Field__000029_Message := (&models.Field{Name: `Message`}).Stage(stage)
	__Field__000030_MessageBackward := (&models.Field{Name: `MessageBackward`}).Stage(stage)
	__Field__000031_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000032_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000033_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000034_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000035_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000036_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000037_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000038_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000039_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000040_Radius := (&models.Field{Name: `Radius`}).Stage(stage)
	__Field__000041_SVG := (&models.Field{Name: `SVG`}).Stage(stage)
	__Field__000042_Speed := (&models.Field{Name: `Speed`}).Stage(stage)
	__Field__000043_StartLat := (&models.Field{Name: `StartLat`}).Stage(stage)
	__Field__000044_StartLng := (&models.Field{Name: `StartLng`}).Stage(stage)
	__Field__000045_TimeOfClick := (&models.Field{Name: `TimeOfClick`}).Stage(stage)
	__Field__000046_UrlTemplate := (&models.Field{Name: `UrlTemplate`}).Stage(stage)
	__Field__000047_VerticalSpeed := (&models.Field{Name: `VerticalSpeed`}).Stage(stage)
	__Field__000048_ZoomControl := (&models.Field{Name: `ZoomControl`}).Stage(stage)
	__Field__000049_ZoomLevel := (&models.Field{Name: `ZoomLevel`}).Stage(stage)
	__Field__000050_ZoomSnap := (&models.Field{Name: `ZoomSnap`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_ColorEnum := (&models.GongEnumShape{Name: `NewDiagram-ColorEnum`}).Stage(stage)
	__GongEnumShape__000001_NewDiagram_DashStyleEnum := (&models.GongEnumShape{Name: `NewDiagram-DashStyleEnum`}).Stage(stage)
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum := (&models.GongEnumShape{Name: `NewDiagram-Start_To_End_Enum`}).Stage(stage)
	__GongEnumShape__000003_NewDiagram_TransmittingEnum := (&models.GongEnumShape{Name: `NewDiagram-TransmittingEnum`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_BACKWARD_END_TO_START := (&models.GongEnumValueEntry{Name: `BACKWARD_END_TO_START`}).Stage(stage)
	__GongEnumValueEntry__000001_BLUE := (&models.GongEnumValueEntry{Name: `BLUE`}).Stage(stage)
	__GongEnumValueEntry__000002_FIVE_TEN := (&models.GongEnumValueEntry{Name: `FIVE_TEN`}).Stage(stage)
	__GongEnumValueEntry__000003_FIVE_TWENTY := (&models.GongEnumValueEntry{Name: `FIVE_TWENTY`}).Stage(stage)
	__GongEnumValueEntry__000004_FORWARD_START_TO_END := (&models.GongEnumValueEntry{Name: `FORWARD_START_TO_END`}).Stage(stage)
	__GongEnumValueEntry__000005_GREEN := (&models.GongEnumValueEntry{Name: `GREEN`}).Stage(stage)
	__GongEnumValueEntry__000006_GREY := (&models.GongEnumValueEntry{Name: `GREY`}).Stage(stage)
	__GongEnumValueEntry__000007_IS_NOT_TRANSMITTING := (&models.GongEnumValueEntry{Name: `IS_NOT_TRANSMITTING`}).Stage(stage)
	__GongEnumValueEntry__000008_IS_TRANSMITTING := (&models.GongEnumValueEntry{Name: `IS_TRANSMITTING`}).Stage(stage)
	__GongEnumValueEntry__000009_LIGHT_BROWN_8D6E63 := (&models.GongEnumValueEntry{Name: `LIGHT_BROWN_8D6E63`}).Stage(stage)
	__GongEnumValueEntry__000010_NONE := (&models.GongEnumValueEntry{Name: `NONE`}).Stage(stage)
	__GongEnumValueEntry__000011_RED := (&models.GongEnumValueEntry{Name: `RED`}).Stage(stage)

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Circle := (&models.GongStructShape{Name: `NewDiagram-Circle`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_DivIcon := (&models.GongStructShape{Name: `NewDiagram-DivIcon`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_LayerGroup := (&models.GongStructShape{Name: `NewDiagram-LayerGroup`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_LayerGroupUse := (&models.GongStructShape{Name: `NewDiagram-LayerGroupUse`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_MapOptions := (&models.GongStructShape{Name: `NewDiagram-MapOptions`}).Stage(stage)
	__GongStructShape__000005_NewDiagram_Marker := (&models.GongStructShape{Name: `NewDiagram-Marker`}).Stage(stage)
	__GongStructShape__000006_NewDiagram_UserClick := (&models.GongStructShape{Name: `NewDiagram-UserClick`}).Stage(stage)
	__GongStructShape__000007_NewDiagram_VLine := (&models.GongStructShape{Name: `NewDiagram-VLine`}).Stage(stage)
	__GongStructShape__000008_NewDiagram_VisualTrack := (&models.GongStructShape{Name: `NewDiagram-VisualTrack`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_DivIcon := (&models.Link{Name: `DivIcon`}).Stage(stage)
	__Link__000001_DivIcon := (&models.Link{Name: `DivIcon`}).Stage(stage)
	__Link__000002_LayerGroup := (&models.Link{Name: `LayerGroup`}).Stage(stage)
	__Link__000003_LayerGroup := (&models.Link{Name: `LayerGroup`}).Stage(stage)
	__Link__000004_LayerGroup := (&models.Link{Name: `LayerGroup`}).Stage(stage)
	__Link__000005_LayerGroup := (&models.Link{Name: `LayerGroup`}).Stage(stage)
	__Link__000006_LayerGroup := (&models.Link{Name: `LayerGroup`}).Stage(stage)
	__Link__000007_LayerGroupUses := (&models.Link{Name: `LayerGroupUses`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Circle := (&models.Position{Name: `Pos-NewDiagram-Circle`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_ColorEnum := (&models.Position{Name: `Pos-NewDiagram-ColorEnum`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_DashStyleEnum := (&models.Position{Name: `Pos-NewDiagram-DashStyleEnum`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_DivIcon := (&models.Position{Name: `Pos-NewDiagram-DivIcon`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_LayerGroup := (&models.Position{Name: `Pos-NewDiagram-LayerGroup`}).Stage(stage)
	__Position__000005_Pos_NewDiagram_LayerGroupUse := (&models.Position{Name: `Pos-NewDiagram-LayerGroupUse`}).Stage(stage)
	__Position__000006_Pos_NewDiagram_MapOptions := (&models.Position{Name: `Pos-NewDiagram-MapOptions`}).Stage(stage)
	__Position__000007_Pos_NewDiagram_Marker := (&models.Position{Name: `Pos-NewDiagram-Marker`}).Stage(stage)
	__Position__000008_Pos_NewDiagram_Start_To_End_Enum := (&models.Position{Name: `Pos-NewDiagram-Start_To_End_Enum`}).Stage(stage)
	__Position__000009_Pos_NewDiagram_TransmittingEnum := (&models.Position{Name: `Pos-NewDiagram-TransmittingEnum`}).Stage(stage)
	__Position__000010_Pos_NewDiagram_UserClick := (&models.Position{Name: `Pos-NewDiagram-UserClick`}).Stage(stage)
	__Position__000011_Pos_NewDiagram_VLine := (&models.Position{Name: `Pos-NewDiagram-VLine`}).Stage(stage)
	__Position__000012_Pos_NewDiagram_VisualTrack := (&models.Position{Name: `Pos-NewDiagram-VisualTrack`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_LayerGroup := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Circle and NewDiagram-LayerGroup`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_LayerGroupUse_and_NewDiagram_LayerGroup := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-LayerGroupUse and NewDiagram-LayerGroup`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_MapOptions_and_NewDiagram_LayerGroupUse := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-MapOptions and NewDiagram-LayerGroupUse`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_DivIcon := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Marker and NewDiagram-DivIcon`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_LayerGroup := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Marker and NewDiagram-LayerGroup`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VLine_and_NewDiagram_LayerGroup := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-VLine and NewDiagram-LayerGroup`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_DivIcon := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-VisualTrack and NewDiagram-DivIcon`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_LayerGroup := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-VisualTrack and NewDiagram-LayerGroup`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Attribution.Name = `Attribution`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.Attribution]
	__Field__000000_Attribution.Identifier = `ref_models.MapOptions.Attribution`
	__Field__000000_Attribution.FieldTypeAsString = ``
	__Field__000000_Attribution.Structname = `MapOptions`
	__Field__000000_Attribution.Fieldtypename = `string`

	// Field values setup
	__Field__000001_AttributionControl.Name = `AttributionControl`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.AttributionControl]
	__Field__000001_AttributionControl.Identifier = `ref_models.MapOptions.AttributionControl`
	__Field__000001_AttributionControl.FieldTypeAsString = ``
	__Field__000001_AttributionControl.Structname = `MapOptions`
	__Field__000001_AttributionControl.Fieldtypename = `bool`

	// Field values setup
	__Field__000002_ColorEnum.Name = `ColorEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.ColorEnum]
	__Field__000002_ColorEnum.Identifier = `ref_models.VLine.ColorEnum`
	__Field__000002_ColorEnum.FieldTypeAsString = ``
	__Field__000002_ColorEnum.Structname = `VLine`
	__Field__000002_ColorEnum.Fieldtypename = `ColorEnum`

	// Field values setup
	__Field__000003_ColorEnum.Name = `ColorEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.ColorEnum]
	__Field__000003_ColorEnum.Identifier = `ref_models.VisualTrack.ColorEnum`
	__Field__000003_ColorEnum.FieldTypeAsString = ``
	__Field__000003_ColorEnum.Structname = `VisualTrack`
	__Field__000003_ColorEnum.Fieldtypename = `ColorEnum`

	// Field values setup
	__Field__000004_ColorEnum.Name = `ColorEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.ColorEnum]
	__Field__000004_ColorEnum.Identifier = `ref_models.Circle.ColorEnum`
	__Field__000004_ColorEnum.FieldTypeAsString = ``
	__Field__000004_ColorEnum.Structname = `Circle`
	__Field__000004_ColorEnum.Fieldtypename = `ColorEnum`

	// Field values setup
	__Field__000005_ColorEnum.Name = `ColorEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Marker.ColorEnum]
	__Field__000005_ColorEnum.Identifier = `ref_models.Marker.ColorEnum`
	__Field__000005_ColorEnum.FieldTypeAsString = ``
	__Field__000005_ColorEnum.Structname = `Marker`
	__Field__000005_ColorEnum.Fieldtypename = `ColorEnum`

	// Field values setup
	__Field__000006_DashStyleEnum.Name = `DashStyleEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.DashStyleEnum]
	__Field__000006_DashStyleEnum.Identifier = `ref_models.Circle.DashStyleEnum`
	__Field__000006_DashStyleEnum.FieldTypeAsString = ``
	__Field__000006_DashStyleEnum.Structname = `Circle`
	__Field__000006_DashStyleEnum.Fieldtypename = `DashStyleEnum`

	// Field values setup
	__Field__000007_DashStyleEnum.Name = `DashStyleEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.DashStyleEnum]
	__Field__000007_DashStyleEnum.Identifier = `ref_models.VLine.DashStyleEnum`
	__Field__000007_DashStyleEnum.FieldTypeAsString = ``
	__Field__000007_DashStyleEnum.Structname = `VLine`
	__Field__000007_DashStyleEnum.Fieldtypename = `DashStyleEnum`

	// Field values setup
	__Field__000008_Display.Name = `Display`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroupUse.Display]
	__Field__000008_Display.Identifier = `ref_models.LayerGroupUse.Display`
	__Field__000008_Display.FieldTypeAsString = ``
	__Field__000008_Display.Structname = `LayerGroupUse`
	__Field__000008_Display.Fieldtypename = `bool`

	// Field values setup
	__Field__000009_DisplayLevelAndSpeed.Name = `DisplayLevelAndSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.DisplayLevelAndSpeed]
	__Field__000009_DisplayLevelAndSpeed.Identifier = `ref_models.VisualTrack.DisplayLevelAndSpeed`
	__Field__000009_DisplayLevelAndSpeed.FieldTypeAsString = ``
	__Field__000009_DisplayLevelAndSpeed.Structname = `VisualTrack`
	__Field__000009_DisplayLevelAndSpeed.Fieldtypename = `bool`

	// Field values setup
	__Field__000010_DisplayName.Name = `DisplayName`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup.DisplayName]
	__Field__000010_DisplayName.Identifier = `ref_models.LayerGroup.DisplayName`
	__Field__000010_DisplayName.FieldTypeAsString = ``
	__Field__000010_DisplayName.Structname = `LayerGroup`
	__Field__000010_DisplayName.Fieldtypename = `string`

	// Field values setup
	__Field__000011_DisplayTrackHistory.Name = `DisplayTrackHistory`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.DisplayTrackHistory]
	__Field__000011_DisplayTrackHistory.Identifier = `ref_models.VisualTrack.DisplayTrackHistory`
	__Field__000011_DisplayTrackHistory.FieldTypeAsString = ``
	__Field__000011_DisplayTrackHistory.Structname = `VisualTrack`
	__Field__000011_DisplayTrackHistory.Fieldtypename = `bool`

	// Field values setup
	__Field__000012_EndLat.Name = `EndLat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.EndLat]
	__Field__000012_EndLat.Identifier = `ref_models.VLine.EndLat`
	__Field__000012_EndLat.FieldTypeAsString = ``
	__Field__000012_EndLat.Structname = `VLine`
	__Field__000012_EndLat.Fieldtypename = `float64`

	// Field values setup
	__Field__000013_EndLng.Name = `EndLng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.EndLng]
	__Field__000013_EndLng.Identifier = `ref_models.VLine.EndLng`
	__Field__000013_EndLng.FieldTypeAsString = ``
	__Field__000013_EndLng.Structname = `VLine`
	__Field__000013_EndLng.Fieldtypename = `float64`

	// Field values setup
	__Field__000014_Heading.Name = `Heading`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.Heading]
	__Field__000014_Heading.Identifier = `ref_models.VisualTrack.Heading`
	__Field__000014_Heading.FieldTypeAsString = ``
	__Field__000014_Heading.Structname = `VisualTrack`
	__Field__000014_Heading.Fieldtypename = `float64`

	// Field values setup
	__Field__000015_IsTransmitting.Name = `IsTransmitting`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.IsTransmitting]
	__Field__000015_IsTransmitting.Identifier = `ref_models.VLine.IsTransmitting`
	__Field__000015_IsTransmitting.FieldTypeAsString = ``
	__Field__000015_IsTransmitting.Structname = `VLine`
	__Field__000015_IsTransmitting.Fieldtypename = `TransmittingEnum`

	// Field values setup
	__Field__000016_IsTransmittingBackward.Name = `IsTransmittingBackward`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.IsTransmittingBackward]
	__Field__000016_IsTransmittingBackward.Identifier = `ref_models.VLine.IsTransmittingBackward`
	__Field__000016_IsTransmittingBackward.FieldTypeAsString = ``
	__Field__000016_IsTransmittingBackward.Structname = `VLine`
	__Field__000016_IsTransmittingBackward.Fieldtypename = `TransmittingEnum`

	// Field values setup
	__Field__000017_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Lat]
	__Field__000017_Lat.Identifier = `ref_models.Circle.Lat`
	__Field__000017_Lat.FieldTypeAsString = ``
	__Field__000017_Lat.Structname = `Circle`
	__Field__000017_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000018_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UserClick.Lat]
	__Field__000018_Lat.Identifier = `ref_models.UserClick.Lat`
	__Field__000018_Lat.FieldTypeAsString = ``
	__Field__000018_Lat.Structname = `UserClick`
	__Field__000018_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000019_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Marker.Lat]
	__Field__000019_Lat.Identifier = `ref_models.Marker.Lat`
	__Field__000019_Lat.FieldTypeAsString = ``
	__Field__000019_Lat.Structname = `Marker`
	__Field__000019_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000020_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.Lat]
	__Field__000020_Lat.Identifier = `ref_models.MapOptions.Lat`
	__Field__000020_Lat.FieldTypeAsString = ``
	__Field__000020_Lat.Structname = `MapOptions`
	__Field__000020_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000021_Lat.Name = `Lat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.Lat]
	__Field__000021_Lat.Identifier = `ref_models.VisualTrack.Lat`
	__Field__000021_Lat.FieldTypeAsString = ``
	__Field__000021_Lat.Structname = `VisualTrack`
	__Field__000021_Lat.Fieldtypename = `float64`

	// Field values setup
	__Field__000022_Level.Name = `Level`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.Level]
	__Field__000022_Level.Identifier = `ref_models.VisualTrack.Level`
	__Field__000022_Level.FieldTypeAsString = ``
	__Field__000022_Level.Structname = `VisualTrack`
	__Field__000022_Level.Fieldtypename = `float64`

	// Field values setup
	__Field__000023_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.Lng]
	__Field__000023_Lng.Identifier = `ref_models.VisualTrack.Lng`
	__Field__000023_Lng.FieldTypeAsString = ``
	__Field__000023_Lng.Structname = `VisualTrack`
	__Field__000023_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000024_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UserClick.Lng]
	__Field__000024_Lng.Identifier = `ref_models.UserClick.Lng`
	__Field__000024_Lng.FieldTypeAsString = ``
	__Field__000024_Lng.Structname = `UserClick`
	__Field__000024_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000025_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Marker.Lng]
	__Field__000025_Lng.Identifier = `ref_models.Marker.Lng`
	__Field__000025_Lng.FieldTypeAsString = ``
	__Field__000025_Lng.Structname = `Marker`
	__Field__000025_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000026_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Lng]
	__Field__000026_Lng.Identifier = `ref_models.Circle.Lng`
	__Field__000026_Lng.FieldTypeAsString = ``
	__Field__000026_Lng.Structname = `Circle`
	__Field__000026_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000027_Lng.Name = `Lng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.Lng]
	__Field__000027_Lng.Identifier = `ref_models.MapOptions.Lng`
	__Field__000027_Lng.FieldTypeAsString = ``
	__Field__000027_Lng.Structname = `MapOptions`
	__Field__000027_Lng.Fieldtypename = `float64`

	// Field values setup
	__Field__000028_MaxZoom.Name = `MaxZoom`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.MaxZoom]
	__Field__000028_MaxZoom.Identifier = `ref_models.MapOptions.MaxZoom`
	__Field__000028_MaxZoom.FieldTypeAsString = ``
	__Field__000028_MaxZoom.Structname = `MapOptions`
	__Field__000028_MaxZoom.Fieldtypename = `int`

	// Field values setup
	__Field__000029_Message.Name = `Message`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.Message]
	__Field__000029_Message.Identifier = `ref_models.VLine.Message`
	__Field__000029_Message.FieldTypeAsString = ``
	__Field__000029_Message.Structname = `VLine`
	__Field__000029_Message.Fieldtypename = `string`

	// Field values setup
	__Field__000030_MessageBackward.Name = `MessageBackward`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.MessageBackward]
	__Field__000030_MessageBackward.Identifier = `ref_models.VLine.MessageBackward`
	__Field__000030_MessageBackward.FieldTypeAsString = ``
	__Field__000030_MessageBackward.Structname = `VLine`
	__Field__000030_MessageBackward.Fieldtypename = `string`

	// Field values setup
	__Field__000031_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DivIcon.Name]
	__Field__000031_Name.Identifier = `ref_models.DivIcon.Name`
	__Field__000031_Name.FieldTypeAsString = ``
	__Field__000031_Name.Structname = `DivIcon`
	__Field__000031_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000032_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup.Name]
	__Field__000032_Name.Identifier = `ref_models.LayerGroup.Name`
	__Field__000032_Name.FieldTypeAsString = ``
	__Field__000032_Name.Structname = `LayerGroup`
	__Field__000032_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000033_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.Name]
	__Field__000033_Name.Identifier = `ref_models.VisualTrack.Name`
	__Field__000033_Name.FieldTypeAsString = ``
	__Field__000033_Name.Structname = `VisualTrack`
	__Field__000033_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000034_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroupUse.Name]
	__Field__000034_Name.Identifier = `ref_models.LayerGroupUse.Name`
	__Field__000034_Name.FieldTypeAsString = ``
	__Field__000034_Name.Structname = `LayerGroupUse`
	__Field__000034_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000035_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Marker.Name]
	__Field__000035_Name.Identifier = `ref_models.Marker.Name`
	__Field__000035_Name.FieldTypeAsString = ``
	__Field__000035_Name.Structname = `Marker`
	__Field__000035_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000036_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.Name]
	__Field__000036_Name.Identifier = `ref_models.VLine.Name`
	__Field__000036_Name.FieldTypeAsString = ``
	__Field__000036_Name.Structname = `VLine`
	__Field__000036_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000037_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.Name]
	__Field__000037_Name.Identifier = `ref_models.MapOptions.Name`
	__Field__000037_Name.FieldTypeAsString = ``
	__Field__000037_Name.Structname = `MapOptions`
	__Field__000037_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000038_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UserClick.Name]
	__Field__000038_Name.Identifier = `ref_models.UserClick.Name`
	__Field__000038_Name.FieldTypeAsString = ``
	__Field__000038_Name.Structname = `UserClick`
	__Field__000038_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000039_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Name]
	__Field__000039_Name.Identifier = `ref_models.Circle.Name`
	__Field__000039_Name.FieldTypeAsString = ``
	__Field__000039_Name.Structname = `Circle`
	__Field__000039_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000040_Radius.Name = `Radius`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Radius]
	__Field__000040_Radius.Identifier = `ref_models.Circle.Radius`
	__Field__000040_Radius.FieldTypeAsString = ``
	__Field__000040_Radius.Structname = `Circle`
	__Field__000040_Radius.Fieldtypename = `float64`

	// Field values setup
	__Field__000041_SVG.Name = `SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DivIcon.SVG]
	__Field__000041_SVG.Identifier = `ref_models.DivIcon.SVG`
	__Field__000041_SVG.FieldTypeAsString = ``
	__Field__000041_SVG.Structname = `DivIcon`
	__Field__000041_SVG.Fieldtypename = `string`

	// Field values setup
	__Field__000042_Speed.Name = `Speed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.Speed]
	__Field__000042_Speed.Identifier = `ref_models.VisualTrack.Speed`
	__Field__000042_Speed.FieldTypeAsString = ``
	__Field__000042_Speed.Structname = `VisualTrack`
	__Field__000042_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000043_StartLat.Name = `StartLat`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.StartLat]
	__Field__000043_StartLat.Identifier = `ref_models.VLine.StartLat`
	__Field__000043_StartLat.FieldTypeAsString = ``
	__Field__000043_StartLat.Structname = `VLine`
	__Field__000043_StartLat.Fieldtypename = `float64`

	// Field values setup
	__Field__000044_StartLng.Name = `StartLng`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.StartLng]
	__Field__000044_StartLng.Identifier = `ref_models.VLine.StartLng`
	__Field__000044_StartLng.FieldTypeAsString = ``
	__Field__000044_StartLng.Structname = `VLine`
	__Field__000044_StartLng.Fieldtypename = `float64`

	// Field values setup
	__Field__000045_TimeOfClick.Name = `TimeOfClick`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UserClick.TimeOfClick]
	__Field__000045_TimeOfClick.Identifier = `ref_models.UserClick.TimeOfClick`
	__Field__000045_TimeOfClick.FieldTypeAsString = ``
	__Field__000045_TimeOfClick.Structname = `UserClick`
	__Field__000045_TimeOfClick.Fieldtypename = `Time`

	// Field values setup
	__Field__000046_UrlTemplate.Name = `UrlTemplate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.UrlTemplate]
	__Field__000046_UrlTemplate.Identifier = `ref_models.MapOptions.UrlTemplate`
	__Field__000046_UrlTemplate.FieldTypeAsString = ``
	__Field__000046_UrlTemplate.Structname = `MapOptions`
	__Field__000046_UrlTemplate.Fieldtypename = `string`

	// Field values setup
	__Field__000047_VerticalSpeed.Name = `VerticalSpeed`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.VerticalSpeed]
	__Field__000047_VerticalSpeed.Identifier = `ref_models.VisualTrack.VerticalSpeed`
	__Field__000047_VerticalSpeed.FieldTypeAsString = ``
	__Field__000047_VerticalSpeed.Structname = `VisualTrack`
	__Field__000047_VerticalSpeed.Fieldtypename = `float64`

	// Field values setup
	__Field__000048_ZoomControl.Name = `ZoomControl`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.ZoomControl]
	__Field__000048_ZoomControl.Identifier = `ref_models.MapOptions.ZoomControl`
	__Field__000048_ZoomControl.FieldTypeAsString = ``
	__Field__000048_ZoomControl.Structname = `MapOptions`
	__Field__000048_ZoomControl.Fieldtypename = `bool`

	// Field values setup
	__Field__000049_ZoomLevel.Name = `ZoomLevel`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.ZoomLevel]
	__Field__000049_ZoomLevel.Identifier = `ref_models.MapOptions.ZoomLevel`
	__Field__000049_ZoomLevel.FieldTypeAsString = ``
	__Field__000049_ZoomLevel.Structname = `MapOptions`
	__Field__000049_ZoomLevel.Fieldtypename = `float64`

	// Field values setup
	__Field__000050_ZoomSnap.Name = `ZoomSnap`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.ZoomSnap]
	__Field__000050_ZoomSnap.Identifier = `ref_models.MapOptions.ZoomSnap`
	__Field__000050_ZoomSnap.FieldTypeAsString = ``
	__Field__000050_ZoomSnap.Structname = `MapOptions`
	__Field__000050_ZoomSnap.Fieldtypename = `int`

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_ColorEnum.Name = `NewDiagram-ColorEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ColorEnum]
	__GongEnumShape__000000_NewDiagram_ColorEnum.Identifier = `ref_models.ColorEnum`
	__GongEnumShape__000000_NewDiagram_ColorEnum.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_ColorEnum.Heigth = 153.000000

	// GongEnumShape values setup
	__GongEnumShape__000001_NewDiagram_DashStyleEnum.Name = `NewDiagram-DashStyleEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DashStyleEnum]
	__GongEnumShape__000001_NewDiagram_DashStyleEnum.Identifier = `ref_models.DashStyleEnum`
	__GongEnumShape__000001_NewDiagram_DashStyleEnum.Width = 240.000000
	__GongEnumShape__000001_NewDiagram_DashStyleEnum.Heigth = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.Name = `NewDiagram-Start_To_End_Enum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Start_To_End_Enum]
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.Identifier = `ref_models.Start_To_End_Enum`
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.Width = 240.000000
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.Heigth = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000003_NewDiagram_TransmittingEnum.Name = `NewDiagram-TransmittingEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TransmittingEnum]
	__GongEnumShape__000003_NewDiagram_TransmittingEnum.Identifier = `ref_models.TransmittingEnum`
	__GongEnumShape__000003_NewDiagram_TransmittingEnum.Width = 240.000000
	__GongEnumShape__000003_NewDiagram_TransmittingEnum.Heigth = 93.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_BACKWARD_END_TO_START.Name = `BACKWARD_END_TO_START`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Start_To_End_Enum.BACKWARD_END_TO_START]
	__GongEnumValueEntry__000000_BACKWARD_END_TO_START.Identifier = `ref_models.Start_To_End_Enum.BACKWARD_END_TO_START`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000001_BLUE.Name = `BLUE`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ColorEnum.BLUE]
	__GongEnumValueEntry__000001_BLUE.Identifier = `ref_models.ColorEnum.BLUE`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000002_FIVE_TEN.Name = `FIVE_TEN`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DashStyleEnum.FIVE_TEN]
	__GongEnumValueEntry__000002_FIVE_TEN.Identifier = `ref_models.DashStyleEnum.FIVE_TEN`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000003_FIVE_TWENTY.Name = `FIVE_TWENTY`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DashStyleEnum.FIVE_TWENTY]
	__GongEnumValueEntry__000003_FIVE_TWENTY.Identifier = `ref_models.DashStyleEnum.FIVE_TWENTY`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000004_FORWARD_START_TO_END.Name = `FORWARD_START_TO_END`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Start_To_End_Enum.FORWARD_START_TO_END]
	__GongEnumValueEntry__000004_FORWARD_START_TO_END.Identifier = `ref_models.Start_To_End_Enum.FORWARD_START_TO_END`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000005_GREEN.Name = `GREEN`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ColorEnum.GREEN]
	__GongEnumValueEntry__000005_GREEN.Identifier = `ref_models.ColorEnum.GREEN`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000006_GREY.Name = `GREY`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ColorEnum.GREY]
	__GongEnumValueEntry__000006_GREY.Identifier = `ref_models.ColorEnum.GREY`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000007_IS_NOT_TRANSMITTING.Name = `IS_NOT_TRANSMITTING`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TransmittingEnum.IS_NOT_TRANSMITTING]
	__GongEnumValueEntry__000007_IS_NOT_TRANSMITTING.Identifier = `ref_models.TransmittingEnum.IS_NOT_TRANSMITTING`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000008_IS_TRANSMITTING.Name = `IS_TRANSMITTING`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TransmittingEnum.IS_TRANSMITTING]
	__GongEnumValueEntry__000008_IS_TRANSMITTING.Identifier = `ref_models.TransmittingEnum.IS_TRANSMITTING`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000009_LIGHT_BROWN_8D6E63.Name = `LIGHT_BROWN_8D6E63`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ColorEnum.LIGHT_BROWN_8D6E63]
	__GongEnumValueEntry__000009_LIGHT_BROWN_8D6E63.Identifier = `ref_models.ColorEnum.LIGHT_BROWN_8D6E63`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000010_NONE.Name = `NONE`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ColorEnum.NONE]
	__GongEnumValueEntry__000010_NONE.Identifier = `ref_models.ColorEnum.NONE`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000011_RED.Name = `RED`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ColorEnum.RED]
	__GongEnumValueEntry__000011_RED.Identifier = `ref_models.ColorEnum.RED`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Circle.Name = `NewDiagram-Circle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle]
	__GongStructShape__000000_NewDiagram_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000000_NewDiagram_Circle.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Circle.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Circle.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Circle.Heigth = 153.000000
	__GongStructShape__000000_NewDiagram_Circle.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_DivIcon.Name = `NewDiagram-DivIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DivIcon]
	__GongStructShape__000001_NewDiagram_DivIcon.Identifier = `ref_models.DivIcon`
	__GongStructShape__000001_NewDiagram_DivIcon.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_DivIcon.NbInstances = 0
	__GongStructShape__000001_NewDiagram_DivIcon.Width = 240.000000
	__GongStructShape__000001_NewDiagram_DivIcon.Heigth = 93.000000
	__GongStructShape__000001_NewDiagram_DivIcon.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_LayerGroup.Name = `NewDiagram-LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup]
	__GongStructShape__000002_NewDiagram_LayerGroup.Identifier = `ref_models.LayerGroup`
	__GongStructShape__000002_NewDiagram_LayerGroup.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_LayerGroup.NbInstances = 0
	__GongStructShape__000002_NewDiagram_LayerGroup.Width = 240.000000
	__GongStructShape__000002_NewDiagram_LayerGroup.Heigth = 93.000000
	__GongStructShape__000002_NewDiagram_LayerGroup.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Name = `NewDiagram-LayerGroupUse`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroupUse]
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Identifier = `ref_models.LayerGroupUse`
	__GongStructShape__000003_NewDiagram_LayerGroupUse.ShowNbInstances = false
	__GongStructShape__000003_NewDiagram_LayerGroupUse.NbInstances = 0
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Width = 240.000000
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Heigth = 93.000000
	__GongStructShape__000003_NewDiagram_LayerGroupUse.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_MapOptions.Name = `NewDiagram-MapOptions`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions]
	__GongStructShape__000004_NewDiagram_MapOptions.Identifier = `ref_models.MapOptions`
	__GongStructShape__000004_NewDiagram_MapOptions.ShowNbInstances = false
	__GongStructShape__000004_NewDiagram_MapOptions.NbInstances = 0
	__GongStructShape__000004_NewDiagram_MapOptions.Width = 240.000000
	__GongStructShape__000004_NewDiagram_MapOptions.Heigth = 213.000000
	__GongStructShape__000004_NewDiagram_MapOptions.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NewDiagram_Marker.Name = `NewDiagram-Marker`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Marker]
	__GongStructShape__000005_NewDiagram_Marker.Identifier = `ref_models.Marker`
	__GongStructShape__000005_NewDiagram_Marker.ShowNbInstances = false
	__GongStructShape__000005_NewDiagram_Marker.NbInstances = 0
	__GongStructShape__000005_NewDiagram_Marker.Width = 240.000000
	__GongStructShape__000005_NewDiagram_Marker.Heigth = 123.000000
	__GongStructShape__000005_NewDiagram_Marker.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NewDiagram_UserClick.Name = `NewDiagram-UserClick`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.UserClick]
	__GongStructShape__000006_NewDiagram_UserClick.Identifier = `ref_models.UserClick`
	__GongStructShape__000006_NewDiagram_UserClick.ShowNbInstances = false
	__GongStructShape__000006_NewDiagram_UserClick.NbInstances = 0
	__GongStructShape__000006_NewDiagram_UserClick.Width = 240.000000
	__GongStructShape__000006_NewDiagram_UserClick.Heigth = 123.000000
	__GongStructShape__000006_NewDiagram_UserClick.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_NewDiagram_VLine.Name = `NewDiagram-VLine`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine]
	__GongStructShape__000007_NewDiagram_VLine.Identifier = `ref_models.VLine`
	__GongStructShape__000007_NewDiagram_VLine.ShowNbInstances = false
	__GongStructShape__000007_NewDiagram_VLine.NbInstances = 0
	__GongStructShape__000007_NewDiagram_VLine.Width = 240.000000
	__GongStructShape__000007_NewDiagram_VLine.Heigth = 228.000000
	__GongStructShape__000007_NewDiagram_VLine.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_NewDiagram_VisualTrack.Name = `NewDiagram-VisualTrack`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack]
	__GongStructShape__000008_NewDiagram_VisualTrack.Identifier = `ref_models.VisualTrack`
	__GongStructShape__000008_NewDiagram_VisualTrack.ShowNbInstances = false
	__GongStructShape__000008_NewDiagram_VisualTrack.NbInstances = 0
	__GongStructShape__000008_NewDiagram_VisualTrack.Width = 240.000000
	__GongStructShape__000008_NewDiagram_VisualTrack.Heigth = 213.000000
	__GongStructShape__000008_NewDiagram_VisualTrack.IsSelected = false

	// Link values setup
	__Link__000000_DivIcon.Name = `DivIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Marker.DivIcon]
	__Link__000000_DivIcon.Identifier = `ref_models.Marker.DivIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DivIcon]
	__Link__000000_DivIcon.Fieldtypename = `ref_models.DivIcon`
	__Link__000000_DivIcon.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_DivIcon.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000001_DivIcon.Name = `DivIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.DivIcon]
	__Link__000001_DivIcon.Identifier = `ref_models.VisualTrack.DivIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DivIcon]
	__Link__000001_DivIcon.Fieldtypename = `ref_models.DivIcon`
	__Link__000001_DivIcon.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_DivIcon.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000002_LayerGroup.Name = `LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroupUse.LayerGroup]
	__Link__000002_LayerGroup.Identifier = `ref_models.LayerGroupUse.LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup]
	__Link__000002_LayerGroup.Fieldtypename = `ref_models.LayerGroup`
	__Link__000002_LayerGroup.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_LayerGroup.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000003_LayerGroup.Name = `LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Marker.LayerGroup]
	__Link__000003_LayerGroup.Identifier = `ref_models.Marker.LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup]
	__Link__000003_LayerGroup.Fieldtypename = `ref_models.LayerGroup`
	__Link__000003_LayerGroup.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_LayerGroup.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000004_LayerGroup.Name = `LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VLine.LayerGroup]
	__Link__000004_LayerGroup.Identifier = `ref_models.VLine.LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup]
	__Link__000004_LayerGroup.Fieldtypename = `ref_models.LayerGroup`
	__Link__000004_LayerGroup.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_LayerGroup.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000005_LayerGroup.Name = `LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.VisualTrack.LayerGroup]
	__Link__000005_LayerGroup.Identifier = `ref_models.VisualTrack.LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup]
	__Link__000005_LayerGroup.Fieldtypename = `ref_models.LayerGroup`
	__Link__000005_LayerGroup.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_LayerGroup.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000006_LayerGroup.Name = `LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.LayerGroup]
	__Link__000006_LayerGroup.Identifier = `ref_models.Circle.LayerGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroup]
	__Link__000006_LayerGroup.Fieldtypename = `ref_models.LayerGroup`
	__Link__000006_LayerGroup.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_LayerGroup.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000007_LayerGroupUses.Name = `LayerGroupUses`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.MapOptions.LayerGroupUses]
	__Link__000007_LayerGroupUses.Identifier = `ref_models.MapOptions.LayerGroupUses`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LayerGroupUse]
	__Link__000007_LayerGroupUses.Fieldtypename = `ref_models.LayerGroupUse`
	__Link__000007_LayerGroupUses.TargetMultiplicity = models.MANY
	__Link__000007_LayerGroupUses.SourceMultiplicity = models.ZERO_ONE

	// Position values setup
	__Position__000000_Pos_NewDiagram_Circle.X = 740.000000
	__Position__000000_Pos_NewDiagram_Circle.Y = 250.000000
	__Position__000000_Pos_NewDiagram_Circle.Name = `Pos-NewDiagram-Circle`

	// Position values setup
	__Position__000001_Pos_NewDiagram_ColorEnum.X = 60.000000
	__Position__000001_Pos_NewDiagram_ColorEnum.Y = 340.000000
	__Position__000001_Pos_NewDiagram_ColorEnum.Name = `Pos-NewDiagram-ColorEnum`

	// Position values setup
	__Position__000002_Pos_NewDiagram_DashStyleEnum.X = 50.000000
	__Position__000002_Pos_NewDiagram_DashStyleEnum.Y = 700.000000
	__Position__000002_Pos_NewDiagram_DashStyleEnum.Name = `Pos-NewDiagram-DashStyleEnum`

	// Position values setup
	__Position__000003_Pos_NewDiagram_DivIcon.X = 340.000000
	__Position__000003_Pos_NewDiagram_DivIcon.Y = 650.000000
	__Position__000003_Pos_NewDiagram_DivIcon.Name = `Pos-NewDiagram-DivIcon`

	// Position values setup
	__Position__000004_Pos_NewDiagram_LayerGroup.X = 960.000000
	__Position__000004_Pos_NewDiagram_LayerGroup.Y = 120.000000
	__Position__000004_Pos_NewDiagram_LayerGroup.Name = `Pos-NewDiagram-LayerGroup`

	// Position values setup
	__Position__000005_Pos_NewDiagram_LayerGroupUse.X = 560.000000
	__Position__000005_Pos_NewDiagram_LayerGroupUse.Y = 120.000000
	__Position__000005_Pos_NewDiagram_LayerGroupUse.Name = `Pos-NewDiagram-LayerGroupUse`

	// Position values setup
	__Position__000006_Pos_NewDiagram_MapOptions.X = 120.000000
	__Position__000006_Pos_NewDiagram_MapOptions.Y = 20.000000
	__Position__000006_Pos_NewDiagram_MapOptions.Name = `Pos-NewDiagram-MapOptions`

	// Position values setup
	__Position__000007_Pos_NewDiagram_Marker.X = 740.000000
	__Position__000007_Pos_NewDiagram_Marker.Y = 420.000000
	__Position__000007_Pos_NewDiagram_Marker.Name = `Pos-NewDiagram-Marker`

	// Position values setup
	__Position__000008_Pos_NewDiagram_Start_To_End_Enum.X = 1120.000000
	__Position__000008_Pos_NewDiagram_Start_To_End_Enum.Y = 690.000000
	__Position__000008_Pos_NewDiagram_Start_To_End_Enum.Name = `Pos-NewDiagram-Start_To_End_Enum`

	// Position values setup
	__Position__000009_Pos_NewDiagram_TransmittingEnum.X = 1110.000000
	__Position__000009_Pos_NewDiagram_TransmittingEnum.Y = 810.000000
	__Position__000009_Pos_NewDiagram_TransmittingEnum.Name = `Pos-NewDiagram-TransmittingEnum`

	// Position values setup
	__Position__000010_Pos_NewDiagram_UserClick.X = 50.000000
	__Position__000010_Pos_NewDiagram_UserClick.Y = 560.000000
	__Position__000010_Pos_NewDiagram_UserClick.Name = `Pos-NewDiagram-UserClick`

	// Position values setup
	__Position__000011_Pos_NewDiagram_VLine.X = 740.000000
	__Position__000011_Pos_NewDiagram_VLine.Y = 560.000000
	__Position__000011_Pos_NewDiagram_VLine.Name = `Pos-NewDiagram-VLine`

	// Position values setup
	__Position__000012_Pos_NewDiagram_VisualTrack.X = 740.000000
	__Position__000012_Pos_NewDiagram_VisualTrack.Y = 820.000000
	__Position__000012_Pos_NewDiagram_VisualTrack.Name = `Pos-NewDiagram-VisualTrack`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_LayerGroup.X = 865.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_LayerGroup.Y = 266.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_LayerGroup.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Circle and NewDiagram-LayerGroup`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_LayerGroupUse_and_NewDiagram_LayerGroup.X = 845.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_LayerGroupUse_and_NewDiagram_LayerGroup.Y = 78.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_LayerGroupUse_and_NewDiagram_LayerGroup.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-LayerGroupUse and NewDiagram-LayerGroup`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_MapOptions_and_NewDiagram_LayerGroupUse.X = 524.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_MapOptions_and_NewDiagram_LayerGroupUse.Y = 35.500000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_MapOptions_and_NewDiagram_LayerGroupUse.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-MapOptions and NewDiagram-LayerGroupUse`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_DivIcon.X = 542.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_DivIcon.Y = 478.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_DivIcon.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Marker and NewDiagram-DivIcon`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_LayerGroup.X = 1030.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_LayerGroup.Y = 376.500000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_LayerGroup.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Marker and NewDiagram-LayerGroup`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VLine_and_NewDiagram_LayerGroup.X = 1210.000000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VLine_and_NewDiagram_LayerGroup.Y = 466.500000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VLine_and_NewDiagram_LayerGroup.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-VLine and NewDiagram-LayerGroup`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_DivIcon.X = 620.000000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_DivIcon.Y = 944.000000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_DivIcon.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-VisualTrack and NewDiagram-DivIcon`

	// Vertice values setup
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_LayerGroup.X = 1205.000000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_LayerGroup.Y = 594.000000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_LayerGroup.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-VisualTrack and NewDiagram-LayerGroup`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Circle)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_LayerGroup)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_LayerGroupUse)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_DivIcon)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_MapOptions)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000005_NewDiagram_Marker)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_UserClick)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000007_NewDiagram_VLine)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000008_NewDiagram_VisualTrack)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000000_NewDiagram_ColorEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000001_NewDiagram_DashStyleEnum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000002_NewDiagram_Start_To_End_Enum)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000003_NewDiagram_TransmittingEnum)
	__GongEnumShape__000000_NewDiagram_ColorEnum.Position = __Position__000001_Pos_NewDiagram_ColorEnum
	__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys, __GongEnumValueEntry__000009_LIGHT_BROWN_8D6E63)
	__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys, __GongEnumValueEntry__000011_RED)
	__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys, __GongEnumValueEntry__000006_GREY)
	__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys, __GongEnumValueEntry__000005_GREEN)
	__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys, __GongEnumValueEntry__000001_BLUE)
	__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ColorEnum.GongEnumValueEntrys, __GongEnumValueEntry__000010_NONE)
	__GongEnumShape__000001_NewDiagram_DashStyleEnum.Position = __Position__000002_Pos_NewDiagram_DashStyleEnum
	__GongEnumShape__000001_NewDiagram_DashStyleEnum.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_DashStyleEnum.GongEnumValueEntrys, __GongEnumValueEntry__000002_FIVE_TEN)
	__GongEnumShape__000001_NewDiagram_DashStyleEnum.GongEnumValueEntrys = append(__GongEnumShape__000001_NewDiagram_DashStyleEnum.GongEnumValueEntrys, __GongEnumValueEntry__000003_FIVE_TWENTY)
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.Position = __Position__000008_Pos_NewDiagram_Start_To_End_Enum
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.GongEnumValueEntrys = append(__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.GongEnumValueEntrys, __GongEnumValueEntry__000004_FORWARD_START_TO_END)
	__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.GongEnumValueEntrys = append(__GongEnumShape__000002_NewDiagram_Start_To_End_Enum.GongEnumValueEntrys, __GongEnumValueEntry__000000_BACKWARD_END_TO_START)
	__GongEnumShape__000003_NewDiagram_TransmittingEnum.Position = __Position__000009_Pos_NewDiagram_TransmittingEnum
	__GongEnumShape__000003_NewDiagram_TransmittingEnum.GongEnumValueEntrys = append(__GongEnumShape__000003_NewDiagram_TransmittingEnum.GongEnumValueEntrys, __GongEnumValueEntry__000008_IS_TRANSMITTING)
	__GongEnumShape__000003_NewDiagram_TransmittingEnum.GongEnumValueEntrys = append(__GongEnumShape__000003_NewDiagram_TransmittingEnum.GongEnumValueEntrys, __GongEnumValueEntry__000007_IS_NOT_TRANSMITTING)
	__GongStructShape__000000_NewDiagram_Circle.Position = __Position__000000_Pos_NewDiagram_Circle
	__GongStructShape__000000_NewDiagram_Circle.Fields = append(__GongStructShape__000000_NewDiagram_Circle.Fields, __Field__000017_Lat)
	__GongStructShape__000000_NewDiagram_Circle.Fields = append(__GongStructShape__000000_NewDiagram_Circle.Fields, __Field__000026_Lng)
	__GongStructShape__000000_NewDiagram_Circle.Fields = append(__GongStructShape__000000_NewDiagram_Circle.Fields, __Field__000039_Name)
	__GongStructShape__000000_NewDiagram_Circle.Fields = append(__GongStructShape__000000_NewDiagram_Circle.Fields, __Field__000040_Radius)
	__GongStructShape__000000_NewDiagram_Circle.Fields = append(__GongStructShape__000000_NewDiagram_Circle.Fields, __Field__000004_ColorEnum)
	__GongStructShape__000000_NewDiagram_Circle.Fields = append(__GongStructShape__000000_NewDiagram_Circle.Fields, __Field__000006_DashStyleEnum)
	__GongStructShape__000000_NewDiagram_Circle.Links = append(__GongStructShape__000000_NewDiagram_Circle.Links, __Link__000006_LayerGroup)
	__GongStructShape__000001_NewDiagram_DivIcon.Position = __Position__000003_Pos_NewDiagram_DivIcon
	__GongStructShape__000001_NewDiagram_DivIcon.Fields = append(__GongStructShape__000001_NewDiagram_DivIcon.Fields, __Field__000031_Name)
	__GongStructShape__000001_NewDiagram_DivIcon.Fields = append(__GongStructShape__000001_NewDiagram_DivIcon.Fields, __Field__000041_SVG)
	__GongStructShape__000002_NewDiagram_LayerGroup.Position = __Position__000004_Pos_NewDiagram_LayerGroup
	__GongStructShape__000002_NewDiagram_LayerGroup.Fields = append(__GongStructShape__000002_NewDiagram_LayerGroup.Fields, __Field__000032_Name)
	__GongStructShape__000002_NewDiagram_LayerGroup.Fields = append(__GongStructShape__000002_NewDiagram_LayerGroup.Fields, __Field__000010_DisplayName)
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Position = __Position__000005_Pos_NewDiagram_LayerGroupUse
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Fields = append(__GongStructShape__000003_NewDiagram_LayerGroupUse.Fields, __Field__000034_Name)
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Fields = append(__GongStructShape__000003_NewDiagram_LayerGroupUse.Fields, __Field__000008_Display)
	__GongStructShape__000003_NewDiagram_LayerGroupUse.Links = append(__GongStructShape__000003_NewDiagram_LayerGroupUse.Links, __Link__000002_LayerGroup)
	__GongStructShape__000004_NewDiagram_MapOptions.Position = __Position__000006_Pos_NewDiagram_MapOptions
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000020_Lat)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000027_Lng)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000037_Name)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000049_ZoomLevel)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000046_UrlTemplate)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000000_Attribution)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000028_MaxZoom)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000048_ZoomControl)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000001_AttributionControl)
	__GongStructShape__000004_NewDiagram_MapOptions.Fields = append(__GongStructShape__000004_NewDiagram_MapOptions.Fields, __Field__000050_ZoomSnap)
	__GongStructShape__000004_NewDiagram_MapOptions.Links = append(__GongStructShape__000004_NewDiagram_MapOptions.Links, __Link__000007_LayerGroupUses)
	__GongStructShape__000005_NewDiagram_Marker.Position = __Position__000007_Pos_NewDiagram_Marker
	__GongStructShape__000005_NewDiagram_Marker.Fields = append(__GongStructShape__000005_NewDiagram_Marker.Fields, __Field__000019_Lat)
	__GongStructShape__000005_NewDiagram_Marker.Fields = append(__GongStructShape__000005_NewDiagram_Marker.Fields, __Field__000025_Lng)
	__GongStructShape__000005_NewDiagram_Marker.Fields = append(__GongStructShape__000005_NewDiagram_Marker.Fields, __Field__000035_Name)
	__GongStructShape__000005_NewDiagram_Marker.Fields = append(__GongStructShape__000005_NewDiagram_Marker.Fields, __Field__000005_ColorEnum)
	__GongStructShape__000005_NewDiagram_Marker.Links = append(__GongStructShape__000005_NewDiagram_Marker.Links, __Link__000000_DivIcon)
	__GongStructShape__000005_NewDiagram_Marker.Links = append(__GongStructShape__000005_NewDiagram_Marker.Links, __Link__000003_LayerGroup)
	__GongStructShape__000006_NewDiagram_UserClick.Position = __Position__000010_Pos_NewDiagram_UserClick
	__GongStructShape__000006_NewDiagram_UserClick.Fields = append(__GongStructShape__000006_NewDiagram_UserClick.Fields, __Field__000038_Name)
	__GongStructShape__000006_NewDiagram_UserClick.Fields = append(__GongStructShape__000006_NewDiagram_UserClick.Fields, __Field__000018_Lat)
	__GongStructShape__000006_NewDiagram_UserClick.Fields = append(__GongStructShape__000006_NewDiagram_UserClick.Fields, __Field__000024_Lng)
	__GongStructShape__000006_NewDiagram_UserClick.Fields = append(__GongStructShape__000006_NewDiagram_UserClick.Fields, __Field__000045_TimeOfClick)
	__GongStructShape__000007_NewDiagram_VLine.Position = __Position__000011_Pos_NewDiagram_VLine
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000043_StartLat)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000044_StartLng)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000012_EndLat)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000013_EndLng)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000036_Name)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000002_ColorEnum)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000007_DashStyleEnum)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000015_IsTransmitting)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000029_Message)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000016_IsTransmittingBackward)
	__GongStructShape__000007_NewDiagram_VLine.Fields = append(__GongStructShape__000007_NewDiagram_VLine.Fields, __Field__000030_MessageBackward)
	__GongStructShape__000007_NewDiagram_VLine.Links = append(__GongStructShape__000007_NewDiagram_VLine.Links, __Link__000004_LayerGroup)
	__GongStructShape__000008_NewDiagram_VisualTrack.Position = __Position__000012_Pos_NewDiagram_VisualTrack
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000021_Lat)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000023_Lng)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000014_Heading)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000022_Level)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000042_Speed)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000047_VerticalSpeed)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000033_Name)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000003_ColorEnum)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000011_DisplayTrackHistory)
	__GongStructShape__000008_NewDiagram_VisualTrack.Fields = append(__GongStructShape__000008_NewDiagram_VisualTrack.Fields, __Field__000009_DisplayLevelAndSpeed)
	__GongStructShape__000008_NewDiagram_VisualTrack.Links = append(__GongStructShape__000008_NewDiagram_VisualTrack.Links, __Link__000001_DivIcon)
	__GongStructShape__000008_NewDiagram_VisualTrack.Links = append(__GongStructShape__000008_NewDiagram_VisualTrack.Links, __Link__000005_LayerGroup)
	__Link__000000_DivIcon.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_DivIcon
	__Link__000001_DivIcon.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_DivIcon
	__Link__000002_LayerGroup.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_LayerGroupUse_and_NewDiagram_LayerGroup
	__Link__000003_LayerGroup.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Marker_and_NewDiagram_LayerGroup
	__Link__000004_LayerGroup.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VLine_and_NewDiagram_LayerGroup
	__Link__000005_LayerGroup.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_VisualTrack_and_NewDiagram_LayerGroup
	__Link__000006_LayerGroup.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_LayerGroup
	__Link__000007_LayerGroupUses.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_MapOptions_and_NewDiagram_LayerGroupUse
}


