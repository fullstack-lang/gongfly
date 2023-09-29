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
var ___dummy__Stage_{{databaseName}} models.StageStruct
var ___dummy__Time_{{databaseName}} time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_{{databaseName}} map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["{{databaseName}}"] = {{databaseName}}Injection
// }

// {{databaseName}}Injection will stage objects of database "{{databaseName}}"
func {{databaseName}}Injection(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}

`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

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
	map_DummyAgent_Identifiers := make(map[*DummyAgent]string)
	_ = map_DummyAgent_Identifiers

	dummyagentOrdered := []*DummyAgent{}
	for dummyagent := range stage.DummyAgents {
		dummyagentOrdered = append(dummyagentOrdered, dummyagent)
	}
	sort.Slice(dummyagentOrdered[:], func(i, j int) bool {
		return dummyagentOrdered[i].Name < dummyagentOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of DummyAgent"
	for idx, dummyagent := range dummyagentOrdered {

		id = generatesIdentifier("DummyAgent", idx, dummyagent.Name)
		map_DummyAgent_Identifiers[dummyagent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DummyAgent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dummyagent.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// DummyAgent values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TechName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dummyagent.TechName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dummyagent.Name))
		initializerStatements += setValueField

	}

	map_Engine_Identifiers := make(map[*Engine]string)
	_ = map_Engine_Identifiers

	engineOrdered := []*Engine{}
	for engine := range stage.Engines {
		engineOrdered = append(engineOrdered, engine)
	}
	sort.Slice(engineOrdered[:], func(i, j int) bool {
		return engineOrdered[i].Name < engineOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Engine"
	for idx, engine := range engineOrdered {

		id = generatesIdentifier("Engine", idx, engine.Name)
		map_Engine_Identifiers[engine] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Engine")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", engine.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Engine values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(engine.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndTime")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(engine.EndTime))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CurrentTime")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(engine.CurrentTime))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SecondsSinceStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", engine.SecondsSinceStart))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fired")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", engine.Fired))
		initializerStatements += setValueField

		if engine.ControlMode != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlMode")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+engine.ControlMode.ToCodeString())
			initializerStatements += setValueField
		}

		if engine.State != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "State")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+engine.State.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Speed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", engine.Speed))
		initializerStatements += setValueField

	}

	map_Event_Identifiers := make(map[*Event]string)
	_ = map_Event_Identifiers

	eventOrdered := []*Event{}
	for event := range stage.Events {
		eventOrdered = append(eventOrdered, event)
	}
	sort.Slice(eventOrdered[:], func(i, j int) bool {
		return eventOrdered[i].Name < eventOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Event"
	for idx, event := range eventOrdered {

		id = generatesIdentifier("Event", idx, event.Name)
		map_Event_Identifiers[event] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Event")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", event.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Event values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(event.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Duration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", event.Duration))
		initializerStatements += setValueField

	}

	map_GongsimCommand_Identifiers := make(map[*GongsimCommand]string)
	_ = map_GongsimCommand_Identifiers

	gongsimcommandOrdered := []*GongsimCommand{}
	for gongsimcommand := range stage.GongsimCommands {
		gongsimcommandOrdered = append(gongsimcommandOrdered, gongsimcommand)
	}
	sort.Slice(gongsimcommandOrdered[:], func(i, j int) bool {
		return gongsimcommandOrdered[i].Name < gongsimcommandOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of GongsimCommand"
	for idx, gongsimcommand := range gongsimcommandOrdered {

		id = generatesIdentifier("GongsimCommand", idx, gongsimcommand.Name)
		map_GongsimCommand_Identifiers[gongsimcommand] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongsimCommand")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongsimcommand.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// GongsimCommand values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongsimcommand.Name))
		initializerStatements += setValueField

		if gongsimcommand.Command != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Command")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongsimcommand.Command.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CommandDate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongsimcommand.CommandDate))
		initializerStatements += setValueField

		if gongsimcommand.SpeedCommandType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SpeedCommandType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongsimcommand.SpeedCommandType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DateSpeedCommand")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongsimcommand.DateSpeedCommand))
		initializerStatements += setValueField

	}

	map_GongsimStatus_Identifiers := make(map[*GongsimStatus]string)
	_ = map_GongsimStatus_Identifiers

	gongsimstatusOrdered := []*GongsimStatus{}
	for gongsimstatus := range stage.GongsimStatuss {
		gongsimstatusOrdered = append(gongsimstatusOrdered, gongsimstatus)
	}
	sort.Slice(gongsimstatusOrdered[:], func(i, j int) bool {
		return gongsimstatusOrdered[i].Name < gongsimstatusOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of GongsimStatus"
	for idx, gongsimstatus := range gongsimstatusOrdered {

		id = generatesIdentifier("GongsimStatus", idx, gongsimstatus.Name)
		map_GongsimStatus_Identifiers[gongsimstatus] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongsimStatus")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongsimstatus.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// GongsimStatus values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongsimstatus.Name))
		initializerStatements += setValueField

		if gongsimstatus.CurrentCommand != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CurrentCommand")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongsimstatus.CurrentCommand.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompletionDate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongsimstatus.CompletionDate))
		initializerStatements += setValueField

		if gongsimstatus.CurrentSpeedCommand != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CurrentSpeedCommand")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongsimstatus.CurrentSpeedCommand.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SpeedCommandCompletionDate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongsimstatus.SpeedCommandCompletionDate))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, dummyagent := range dummyagentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DummyAgent", idx, dummyagent.Name)
		map_DummyAgent_Identifiers[dummyagent] = id

		// Initialisation of values
	}

	for idx, engine := range engineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Engine", idx, engine.Name)
		map_Engine_Identifiers[engine] = id

		// Initialisation of values
	}

	for idx, event := range eventOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Event", idx, event.Name)
		map_Event_Identifiers[event] = id

		// Initialisation of values
	}

	for idx, gongsimcommand := range gongsimcommandOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongsimCommand", idx, gongsimcommand.Name)
		map_GongsimCommand_Identifiers[gongsimcommand] = id

		// Initialisation of values
		if gongsimcommand.Engine != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Engine")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Engine_Identifiers[gongsimcommand.Engine])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, gongsimstatus := range gongsimstatusOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongsimStatus", idx, gongsimstatus.Name)
		map_GongsimStatus_Identifiers[gongsimstatus] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar ___dummy__%s_%s %s.StageStruct",
				stage.MetaPackageImportAlias,
				strings.ReplaceAll(filepath.Base(name), ".go", ""),
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

		res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
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
