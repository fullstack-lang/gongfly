// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Circle_Identifiers := make(map[*Circle]string)
	_ = map_Circle_Identifiers

	circleOrdered := []*Circle{}
	for circle := range stage.Circles {
		circleOrdered = append(circleOrdered, circle)
	}
	sort.Slice(circleOrdered[:], func(i, j int) bool {
		return circleOrdered[i].Name < circleOrdered[j].Name
	})
	if len(circleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, circle := range circleOrdered {

		id = generatesIdentifier("Circle", idx, circle.Name)
		map_Circle_Identifiers[circle] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Circle")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", circle.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.Lat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lng")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.Lng))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Radius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.Radius))
		initializerStatements += setValueField

		if circle.ColorEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ColorEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+circle.ColorEnum.ToCodeString())
			initializerStatements += setValueField
		}

		if circle.DashStyleEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DashStyleEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+circle.DashStyleEnum.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_DivIcon_Identifiers := make(map[*DivIcon]string)
	_ = map_DivIcon_Identifiers

	diviconOrdered := []*DivIcon{}
	for divicon := range stage.DivIcons {
		diviconOrdered = append(diviconOrdered, divicon)
	}
	sort.Slice(diviconOrdered[:], func(i, j int) bool {
		return diviconOrdered[i].Name < diviconOrdered[j].Name
	})
	if len(diviconOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, divicon := range diviconOrdered {

		id = generatesIdentifier("DivIcon", idx, divicon.Name)
		map_DivIcon_Identifiers[divicon] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DivIcon")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", divicon.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(divicon.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(divicon.SVG))
		initializerStatements += setValueField

	}

	map_LayerGroup_Identifiers := make(map[*LayerGroup]string)
	_ = map_LayerGroup_Identifiers

	layergroupOrdered := []*LayerGroup{}
	for layergroup := range stage.LayerGroups {
		layergroupOrdered = append(layergroupOrdered, layergroup)
	}
	sort.Slice(layergroupOrdered[:], func(i, j int) bool {
		return layergroupOrdered[i].Name < layergroupOrdered[j].Name
	})
	if len(layergroupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, layergroup := range layergroupOrdered {

		id = generatesIdentifier("LayerGroup", idx, layergroup.Name)
		map_LayerGroup_Identifiers[layergroup] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LayerGroup")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", layergroup.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(layergroup.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DisplayName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(layergroup.DisplayName))
		initializerStatements += setValueField

	}

	map_LayerGroupUse_Identifiers := make(map[*LayerGroupUse]string)
	_ = map_LayerGroupUse_Identifiers

	layergroupuseOrdered := []*LayerGroupUse{}
	for layergroupuse := range stage.LayerGroupUses {
		layergroupuseOrdered = append(layergroupuseOrdered, layergroupuse)
	}
	sort.Slice(layergroupuseOrdered[:], func(i, j int) bool {
		return layergroupuseOrdered[i].Name < layergroupuseOrdered[j].Name
	})
	if len(layergroupuseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, layergroupuse := range layergroupuseOrdered {

		id = generatesIdentifier("LayerGroupUse", idx, layergroupuse.Name)
		map_LayerGroupUse_Identifiers[layergroupuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LayerGroupUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", layergroupuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(layergroupuse.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", layergroupuse.IsDisplayed))
		initializerStatements += setValueField

	}

	map_MapOptions_Identifiers := make(map[*MapOptions]string)
	_ = map_MapOptions_Identifiers

	mapoptionsOrdered := []*MapOptions{}
	for mapoptions := range stage.MapOptionss {
		mapoptionsOrdered = append(mapoptionsOrdered, mapoptions)
	}
	sort.Slice(mapoptionsOrdered[:], func(i, j int) bool {
		return mapoptionsOrdered[i].Name < mapoptionsOrdered[j].Name
	})
	if len(mapoptionsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, mapoptions := range mapoptionsOrdered {

		id = generatesIdentifier("MapOptions", idx, mapoptions.Name)
		map_MapOptions_Identifiers[mapoptions] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MapOptions")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", mapoptions.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", mapoptions.Lat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lng")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", mapoptions.Lng))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mapoptions.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ZoomLevel")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", mapoptions.ZoomLevel))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "UrlTemplate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mapoptions.UrlTemplate))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Attribution")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mapoptions.Attribution))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxZoom")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", mapoptions.MaxZoom))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ZoomControl")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", mapoptions.ZoomControl))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AttributionControl")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", mapoptions.AttributionControl))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ZoomSnap")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", mapoptions.ZoomSnap))
		initializerStatements += setValueField

	}

	map_Marker_Identifiers := make(map[*Marker]string)
	_ = map_Marker_Identifiers

	markerOrdered := []*Marker{}
	for marker := range stage.Markers {
		markerOrdered = append(markerOrdered, marker)
	}
	sort.Slice(markerOrdered[:], func(i, j int) bool {
		return markerOrdered[i].Name < markerOrdered[j].Name
	})
	if len(markerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, marker := range markerOrdered {

		id = generatesIdentifier("Marker", idx, marker.Name)
		map_Marker_Identifiers[marker] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Marker")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", marker.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", marker.Lat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lng")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", marker.Lng))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(marker.Name))
		initializerStatements += setValueField

		if marker.ColorEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ColorEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+marker.ColorEnum.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_UserClick_Identifiers := make(map[*UserClick]string)
	_ = map_UserClick_Identifiers

	userclickOrdered := []*UserClick{}
	for userclick := range stage.UserClicks {
		userclickOrdered = append(userclickOrdered, userclick)
	}
	sort.Slice(userclickOrdered[:], func(i, j int) bool {
		return userclickOrdered[i].Name < userclickOrdered[j].Name
	})
	if len(userclickOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, userclick := range userclickOrdered {

		id = generatesIdentifier("UserClick", idx, userclick.Name)
		map_UserClick_Identifiers[userclick] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UserClick")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", userclick.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(userclick.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", userclick.Lat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lng")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", userclick.Lng))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeOfClick")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", userclick.TimeOfClick.String())
		initializerStatements += setValueField

	}

	map_VLine_Identifiers := make(map[*VLine]string)
	_ = map_VLine_Identifiers

	vlineOrdered := []*VLine{}
	for vline := range stage.VLines {
		vlineOrdered = append(vlineOrdered, vline)
	}
	sort.Slice(vlineOrdered[:], func(i, j int) bool {
		return vlineOrdered[i].Name < vlineOrdered[j].Name
	})
	if len(vlineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, vline := range vlineOrdered {

		id = generatesIdentifier("VLine", idx, vline.Name)
		map_VLine_Identifiers[vline] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VLine")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", vline.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartLat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", vline.StartLat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartLng")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", vline.StartLng))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndLat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", vline.EndLat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndLng")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", vline.EndLng))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(vline.Name))
		initializerStatements += setValueField

		if vline.ColorEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ColorEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+vline.ColorEnum.ToCodeString())
			initializerStatements += setValueField
		}

		if vline.DashStyleEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DashStyleEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+vline.DashStyleEnum.ToCodeString())
			initializerStatements += setValueField
		}

		if vline.IsTransmitting != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsTransmitting")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+vline.IsTransmitting.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Message")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(vline.Message))
		initializerStatements += setValueField

		if vline.IsTransmittingBackward != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsTransmittingBackward")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+vline.IsTransmittingBackward.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MessageBackward")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(vline.MessageBackward))
		initializerStatements += setValueField

	}

	map_VisualTrack_Identifiers := make(map[*VisualTrack]string)
	_ = map_VisualTrack_Identifiers

	visualtrackOrdered := []*VisualTrack{}
	for visualtrack := range stage.VisualTracks {
		visualtrackOrdered = append(visualtrackOrdered, visualtrack)
	}
	sort.Slice(visualtrackOrdered[:], func(i, j int) bool {
		return visualtrackOrdered[i].Name < visualtrackOrdered[j].Name
	})
	if len(visualtrackOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, visualtrack := range visualtrackOrdered {

		id = generatesIdentifier("VisualTrack", idx, visualtrack.Name)
		map_VisualTrack_Identifiers[visualtrack] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VisualTrack")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", visualtrack.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", visualtrack.Lat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lng")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", visualtrack.Lng))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Heading")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", visualtrack.Heading))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Level")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", visualtrack.Level))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Speed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", visualtrack.Speed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "VerticalSpeed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", visualtrack.VerticalSpeed))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(visualtrack.Name))
		initializerStatements += setValueField

		if visualtrack.ColorEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ColorEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+visualtrack.ColorEnum.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DisplayTrackHistory")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", visualtrack.DisplayTrackHistory))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DisplayLevelAndSpeed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", visualtrack.DisplayLevelAndSpeed))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, circle := range circleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Circle", idx, circle.Name)
		map_Circle_Identifiers[circle] = id

		// Initialisation of values
		if circle.LayerGroup != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LayerGroup")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LayerGroup_Identifiers[circle.LayerGroup])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, divicon := range diviconOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DivIcon", idx, divicon.Name)
		map_DivIcon_Identifiers[divicon] = id

		// Initialisation of values
	}

	for idx, layergroup := range layergroupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LayerGroup", idx, layergroup.Name)
		map_LayerGroup_Identifiers[layergroup] = id

		// Initialisation of values
	}

	for idx, layergroupuse := range layergroupuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LayerGroupUse", idx, layergroupuse.Name)
		map_LayerGroupUse_Identifiers[layergroupuse] = id

		// Initialisation of values
		if layergroupuse.LayerGroup != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LayerGroup")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LayerGroup_Identifiers[layergroupuse.LayerGroup])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, mapoptions := range mapoptionsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MapOptions", idx, mapoptions.Name)
		map_MapOptions_Identifiers[mapoptions] = id

		// Initialisation of values
		for _, _layergroupuse := range mapoptions.LayerGroupUses {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LayerGroupUses")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LayerGroupUse_Identifiers[_layergroupuse])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, marker := range markerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Marker", idx, marker.Name)
		map_Marker_Identifiers[marker] = id

		// Initialisation of values
		if marker.LayerGroup != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LayerGroup")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LayerGroup_Identifiers[marker.LayerGroup])
			pointersInitializesStatements += setPointerField
		}

		if marker.DivIcon != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DivIcon")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DivIcon_Identifiers[marker.DivIcon])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, userclick := range userclickOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("UserClick", idx, userclick.Name)
		map_UserClick_Identifiers[userclick] = id

		// Initialisation of values
	}

	for idx, vline := range vlineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("VLine", idx, vline.Name)
		map_VLine_Identifiers[vline] = id

		// Initialisation of values
		if vline.LayerGroup != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LayerGroup")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LayerGroup_Identifiers[vline.LayerGroup])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, visualtrack := range visualtrackOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("VisualTrack", idx, visualtrack.Name)
		map_VisualTrack_Identifiers[visualtrack] = id

		// Initialisation of values
		if visualtrack.LayerGroup != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LayerGroup")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LayerGroup_Identifiers[visualtrack.LayerGroup])
			pointersInitializesStatements += setPointerField
		}

		if visualtrack.DivIcon != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DivIcon")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DivIcon_Identifiers[visualtrack.DivIcon])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
