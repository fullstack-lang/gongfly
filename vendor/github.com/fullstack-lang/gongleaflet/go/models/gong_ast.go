package models

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(pathToFile string) error {

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "Injection")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName := UnmarshallGongstructStaging(assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_CheckoutScheduler = make(map[string]*CheckoutScheduler)
var __gong__map_Circle = make(map[string]*Circle)
var __gong__map_DivIcon = make(map[string]*DivIcon)
var __gong__map_LayerGroup = make(map[string]*LayerGroup)
var __gong__map_LayerGroupUse = make(map[string]*LayerGroupUse)
var __gong__map_MapOptions = make(map[string]*MapOptions)
var __gong__map_Marker = make(map[string]*Marker)
var __gong__map_UserClick = make(map[string]*UserClick)
var __gong__map_VLine = make(map[string]*VLine)
var __gong__map_VisualTrack = make(map[string]*VisualTrack)

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {
	astCoordinate := "\tAssignStmt: "
	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "CheckoutScheduler":
										instanceCheckoutScheduler := (&CheckoutScheduler{Name: instanceName}).Stage()
										instance = any(instanceCheckoutScheduler)
										__gong__map_CheckoutScheduler[identifier] = instanceCheckoutScheduler
									case "Circle":
										instanceCircle := (&Circle{Name: instanceName}).Stage()
										instance = any(instanceCircle)
										__gong__map_Circle[identifier] = instanceCircle
									case "DivIcon":
										instanceDivIcon := (&DivIcon{Name: instanceName}).Stage()
										instance = any(instanceDivIcon)
										__gong__map_DivIcon[identifier] = instanceDivIcon
									case "LayerGroup":
										instanceLayerGroup := (&LayerGroup{Name: instanceName}).Stage()
										instance = any(instanceLayerGroup)
										__gong__map_LayerGroup[identifier] = instanceLayerGroup
									case "LayerGroupUse":
										instanceLayerGroupUse := (&LayerGroupUse{Name: instanceName}).Stage()
										instance = any(instanceLayerGroupUse)
										__gong__map_LayerGroupUse[identifier] = instanceLayerGroupUse
									case "MapOptions":
										instanceMapOptions := (&MapOptions{Name: instanceName}).Stage()
										instance = any(instanceMapOptions)
										__gong__map_MapOptions[identifier] = instanceMapOptions
									case "Marker":
										instanceMarker := (&Marker{Name: instanceName}).Stage()
										instance = any(instanceMarker)
										__gong__map_Marker[identifier] = instanceMarker
									case "UserClick":
										instanceUserClick := (&UserClick{Name: instanceName}).Stage()
										instance = any(instanceUserClick)
										__gong__map_UserClick[identifier] = instanceUserClick
									case "VLine":
										instanceVLine := (&VLine{Name: instanceName}).Stage()
										instance = any(instanceVLine)
										__gong__map_VLine[identifier] = instanceVLine
									case "VisualTrack":
										instanceVisualTrack := (&VisualTrack{Name: instanceName}).Stage()
										instance = any(instanceVisualTrack)
										__gong__map_VisualTrack[identifier] = instanceVisualTrack
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "CheckoutScheduler":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Circle":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DivIcon":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LayerGroup":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LayerGroupUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MapOptions":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Marker":
							switch fieldName {
							// insertion point for date assign code
							}
						case "UserClick":
							switch fieldName {
							// insertion point for date assign code
							case "TimeOfClick":
								__gong__map_UserClick[identifier].TimeOfClick, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "VLine":
							switch fieldName {
							// insertion point for date assign code
							}
						case "VisualTrack":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "CheckoutScheduler":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Circle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DivIcon":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "LayerGroup":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "LayerGroupUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "MapOptions":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "LayerGroupUses":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_LayerGroupUse[targetIdentifier]
							__gong__map_MapOptions[identifier].LayerGroupUses =
								append(__gong__map_MapOptions[identifier].LayerGroupUses, target)
						}
					case "Marker":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "UserClick":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "VLine":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "VisualTrack":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit:
			// assignment to string field
			basicLit := expr
			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "CheckoutScheduler":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CheckoutScheduler[identifier].Name = fielValue
				case "NbUpdatesFromFront":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_CheckoutScheduler[identifier].NbUpdatesFromFront = int(fielValue)
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				case "Lat":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].Lat = fielValue
				case "Lng":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].Lng = fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Name = fielValue
				case "Radius":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].Radius = fielValue
				}
			case "DivIcon":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DivIcon[identifier].Name = fielValue
				case "SVG":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DivIcon[identifier].SVG = fielValue
				}
			case "LayerGroup":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LayerGroup[identifier].Name = fielValue
				case "DisplayName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LayerGroup[identifier].DisplayName = fielValue
				}
			case "LayerGroupUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LayerGroupUse[identifier].Name = fielValue
				}
			case "MapOptions":
				switch fieldName {
				// insertion point for field dependant code
				case "Lat":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_MapOptions[identifier].Lat = fielValue
				case "Lng":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_MapOptions[identifier].Lng = fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MapOptions[identifier].Name = fielValue
				case "ZoomLevel":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_MapOptions[identifier].ZoomLevel = fielValue
				case "UrlTemplate":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MapOptions[identifier].UrlTemplate = fielValue
				case "Attribution":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MapOptions[identifier].Attribution = fielValue
				case "MaxZoom":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_MapOptions[identifier].MaxZoom = int(fielValue)
				case "ZoomSnap":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_MapOptions[identifier].ZoomSnap = int(fielValue)
				}
			case "Marker":
				switch fieldName {
				// insertion point for field dependant code
				case "Lat":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Marker[identifier].Lat = fielValue
				case "Lng":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Marker[identifier].Lng = fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Marker[identifier].Name = fielValue
				}
			case "UserClick":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_UserClick[identifier].Name = fielValue
				case "Lat":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_UserClick[identifier].Lat = fielValue
				case "Lng":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_UserClick[identifier].Lng = fielValue
				}
			case "VLine":
				switch fieldName {
				// insertion point for field dependant code
				case "StartLat":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VLine[identifier].StartLat = fielValue
				case "StartLng":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VLine[identifier].StartLng = fielValue
				case "EndLat":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VLine[identifier].EndLat = fielValue
				case "EndLng":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VLine[identifier].EndLng = fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VLine[identifier].Name = fielValue
				case "Message":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VLine[identifier].Message = fielValue
				case "MessageBackward":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VLine[identifier].MessageBackward = fielValue
				}
			case "VisualTrack":
				switch fieldName {
				// insertion point for field dependant code
				case "Lat":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].Lat = fielValue
				case "Lng":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].Lng = fielValue
				case "Heading":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].Heading = fielValue
				case "Level":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].Level = fielValue
				case "Speed":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].Speed = fielValue
				case "VerticalSpeed":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].VerticalSpeed = fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VisualTrack[identifier].Name = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "CheckoutScheduler":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				case "LayerGroup":
					targetIdentifier := ident.Name
					__gong__map_Circle[identifier].LayerGroup = __gong__map_LayerGroup[targetIdentifier]
				}
			case "DivIcon":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "LayerGroup":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "LayerGroupUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Display":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_LayerGroupUse[identifier].Display = fielValue
				case "LayerGroup":
					targetIdentifier := ident.Name
					__gong__map_LayerGroupUse[identifier].LayerGroup = __gong__map_LayerGroup[targetIdentifier]
				}
			case "MapOptions":
				switch fieldName {
				// insertion point for field dependant code
				case "ZoomControl":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_MapOptions[identifier].ZoomControl = fielValue
				case "AttributionControl":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_MapOptions[identifier].AttributionControl = fielValue
				}
			case "Marker":
				switch fieldName {
				// insertion point for field dependant code
				case "LayerGroup":
					targetIdentifier := ident.Name
					__gong__map_Marker[identifier].LayerGroup = __gong__map_LayerGroup[targetIdentifier]
				case "DivIcon":
					targetIdentifier := ident.Name
					__gong__map_Marker[identifier].DivIcon = __gong__map_DivIcon[targetIdentifier]
				}
			case "UserClick":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "VLine":
				switch fieldName {
				// insertion point for field dependant code
				case "LayerGroup":
					targetIdentifier := ident.Name
					__gong__map_VLine[identifier].LayerGroup = __gong__map_LayerGroup[targetIdentifier]
				}
			case "VisualTrack":
				switch fieldName {
				// insertion point for field dependant code
				case "LayerGroup":
					targetIdentifier := ident.Name
					__gong__map_VisualTrack[identifier].LayerGroup = __gong__map_LayerGroup[targetIdentifier]
				case "DivIcon":
					targetIdentifier := ident.Name
					__gong__map_VisualTrack[identifier].DivIcon = __gong__map_DivIcon[targetIdentifier]
				case "DisplayTrackHistory":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].DisplayTrackHistory = fielValue
				case "DisplayLevelAndSpeed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VisualTrack[identifier].DisplayLevelAndSpeed = fielValue
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "CheckoutScheduler":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Circle":
					switch fieldName {
					// insertion point for enum assign code
					case "ColorEnum":
						var val ColorEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Circle[identifier].ColorEnum = ColorEnum(val)
					case "DashStyleEnum":
						var val DashStyleEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Circle[identifier].DashStyleEnum = DashStyleEnum(val)
					}
				case "DivIcon":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "LayerGroup":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "LayerGroupUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MapOptions":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Marker":
					switch fieldName {
					// insertion point for enum assign code
					case "ColorEnum":
						var val ColorEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Marker[identifier].ColorEnum = ColorEnum(val)
					}
				case "UserClick":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "VLine":
					switch fieldName {
					// insertion point for enum assign code
					case "ColorEnum":
						var val ColorEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_VLine[identifier].ColorEnum = ColorEnum(val)
					case "DashStyleEnum":
						var val DashStyleEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_VLine[identifier].DashStyleEnum = DashStyleEnum(val)
					case "IsTransmitting":
						var val TransmittingEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_VLine[identifier].IsTransmitting = TransmittingEnum(val)
					case "IsTransmittingBackward":
						var val TransmittingEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_VLine[identifier].IsTransmittingBackward = TransmittingEnum(val)
					}
				case "VisualTrack":
					switch fieldName {
					// insertion point for enum assign code
					case "ColorEnum":
						var val ColorEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_VisualTrack[identifier].ColorEnum = ColorEnum(val)
					}
				}
			}
		}
	}
	return
}
