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
	case DummyAgent:
		fieldCoder := DummyAgent{}
		// insertion point for field dependant code
		fieldCoder.TechName = "0"
		fieldCoder.Name = "1"
		return (any)(fieldCoder).(Type)
	case Engine:
		fieldCoder := Engine{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.EndTime = "1"
		fieldCoder.CurrentTime = "2"
		fieldCoder.SecondsSinceStart = 3.000000
		fieldCoder.Fired = 4
		fieldCoder.ControlMode = "5"
		fieldCoder.State = "6"
		fieldCoder.Speed = 7.000000
		return (any)(fieldCoder).(Type)
	case Event:
		fieldCoder := Event{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Duration = 1
		return (any)(fieldCoder).(Type)
	case GongsimCommand:
		fieldCoder := GongsimCommand{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Command = "1"
		fieldCoder.CommandDate = "2"
		fieldCoder.SpeedCommandType = "3"
		fieldCoder.DateSpeedCommand = "4"
		return (any)(fieldCoder).(Type)
	case GongsimStatus:
		fieldCoder := GongsimStatus{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.CurrentCommand = "1"
		fieldCoder.CompletionDate = "2"
		fieldCoder.CurrentSpeedCommand = "3"
		fieldCoder.SpeedCommandCompletionDate = "4"
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *DummyAgent | []*DummyAgent | *Engine | []*Engine | *Event | []*Event | *GongsimCommand | []*GongsimCommand | *GongsimStatus | []*GongsimStatus
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
	case *DummyAgent:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "TechName"
			}
			if field == "1" {
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
	case *Engine:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "EndTime"
			}
			if field == "2" {
				return "CurrentTime"
			}
			if field == "5" {
				return "ControlMode"
			}
			if field == "6" {
				return "State"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 4 {
				return "Fired"
			}
		case float64:
			// insertion point for field dependant name
			if field == 3.000000 {
				return "SecondsSinceStart"
			}
			if field == 7.000000 {
				return "Speed"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Event:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "Duration"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *GongsimCommand:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Command"
			}
			if field == "2" {
				return "CommandDate"
			}
			if field == "3" {
				return "SpeedCommandType"
			}
			if field == "4" {
				return "DateSpeedCommand"
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
	case *GongsimStatus:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "CurrentCommand"
			}
			if field == "2" {
				return "CompletionDate"
			}
			if field == "3" {
				return "CurrentSpeedCommand"
			}
			if field == "4" {
				return "SpeedCommandCompletionDate"
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
