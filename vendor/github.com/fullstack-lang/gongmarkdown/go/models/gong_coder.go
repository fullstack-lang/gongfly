package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case AnotherDummyData:
		fieldCoder := AnotherDummyData{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	case Cell:
		fieldCoder := Cell{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	case DummyData:
		fieldCoder := DummyData{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.DummyString = "1"
		fieldCoder.DummyInt = 2
		fieldCoder.DummyFloat = 3.000000
		fieldCoder.DummyBool = false
		fieldCoder.DummyEnumString = "5"
		fieldCoder.DummyEnumInt = 6
		fieldCoder.DummyTime = time.Date(7, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.DummyDuration = 8
		return (any)(fieldCoder).(Type)
	case Element:
		fieldCoder := Element{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Content = "1"
		fieldCoder.Type = "2"
		return (any)(fieldCoder).(Type)
	case MarkdownContent:
		fieldCoder := MarkdownContent{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Content = "1"
		return (any)(fieldCoder).(Type)
	case Row:
		fieldCoder := Row{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *AnotherDummyData | []*AnotherDummyData | *Cell | []*Cell | *DummyData | []*DummyData | *Element | []*Element | *MarkdownContent | []*MarkdownContent | *Row | []*Row
}

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case *AnotherDummyData:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Cell:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *DummyData:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "DummyString"
			}
			if field == "5" {
				return "DummyEnumString"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 2 {
				return "DummyInt"
			}
			if field == 6 {
				return "DummyEnumInt"
			}
			if field == 8 {
				return "DummyDuration"
			}
		case float64:
			// insertion point for field dependant name
			if field == 3.000000 {
				return "DummyFloat"
			}
		case time.Time:
			// insertion point for field dependant name
			if field == time.Date(7, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "DummyTime"
			}
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "DummyBool"
			}
		}
	case *Element:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Content"
			}
			if field == "2" {
				return "Type"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *MarkdownContent:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Content"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Row:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
