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
	case CivilianAirport:
		fieldCoder := CivilianAirport{}
		// insertion point for field dependant code
		fieldCoder.Lat = 0.000000
		fieldCoder.Lng = 1.000000
		fieldCoder.Name = "2"
		return (any)(fieldCoder).(Type)
	case Liner:
		fieldCoder := Liner{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Lat = 1.000000
		fieldCoder.Lng = 2.000000
		fieldCoder.Heading = 3.000000
		fieldCoder.Level = 4.000000
		fieldCoder.Speed = 5.000000
		fieldCoder.State = "6"
		fieldCoder.TargetHeading = 7.000000
		fieldCoder.TargetLocationLat = 8.000000
		fieldCoder.TargetLocationLng = 9.000000
		fieldCoder.DistanceToTarget = 10.000000
		fieldCoder.MaxRotationalSpeed = 11.000000
		fieldCoder.VerticalSpeed = 12.000000
		fieldCoder.Timestampstring = "13"
		return (any)(fieldCoder).(Type)
	case Message:
		fieldCoder := Message{}
		// insertion point for field dependant code
		fieldCoder.Lat = 0.000000
		fieldCoder.Lng = 1.000000
		fieldCoder.Heading = 2.000000
		fieldCoder.Level = 3.000000
		fieldCoder.Speed = 4.000000
		fieldCoder.State = "5"
		fieldCoder.Name = "6"
		fieldCoder.TargetLocationLat = 7.000000
		fieldCoder.TargetLocationLng = 8.000000
		fieldCoder.DistanceToTarget = 9.000000
		fieldCoder.Timestampstring = "10"
		fieldCoder.DurationSinceSimulationStart = 11
		fieldCoder.Timestampstartstring = "12"
		fieldCoder.Source = "13"
		fieldCoder.Destination = "14"
		fieldCoder.Content = "15"
		fieldCoder.About_string = "16"
		fieldCoder.Display = false
		return (any)(fieldCoder).(Type)
	case OpsLine:
		fieldCoder := OpsLine{}
		// insertion point for field dependant code
		fieldCoder.IsTransmitting = false
		fieldCoder.TransmissionMessage = "1"
		fieldCoder.IsTransmittingBackward = true
		fieldCoder.TransmissionMessageBackward = "3"
		fieldCoder.State = "5"
		fieldCoder.Name = "6"
		return (any)(fieldCoder).(Type)
	case Radar:
		fieldCoder := Radar{}
		// insertion point for field dependant code
		fieldCoder.State = "0"
		fieldCoder.Name = "1"
		fieldCoder.Lat = 2.000000
		fieldCoder.Lng = 3.000000
		fieldCoder.Range = 4.000000
		return (any)(fieldCoder).(Type)
	case Satellite:
		fieldCoder := Satellite{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Line1 = "1"
		fieldCoder.Line2 = "2"
		fieldCoder.Lat = 3.000000
		fieldCoder.Lng = 4.000000
		fieldCoder.Heading = 5.000000
		fieldCoder.Level = 6.000000
		fieldCoder.Speed = 7.000000
		fieldCoder.VerticalSpeed = 8.000000
		fieldCoder.Timestampstring = "9"
		return (any)(fieldCoder).(Type)
	case Scenario:
		fieldCoder := Scenario{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Lat = 1.000000
		fieldCoder.Lng = 2.000000
		fieldCoder.ZoomLevel = 3.000000
		fieldCoder.MessageVisualSpeed = 4.000000
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *CivilianAirport | []*CivilianAirport | *Liner | []*Liner | *Message | []*Message | *OpsLine | []*OpsLine | *Radar | []*Radar | *Satellite | []*Satellite | *Scenario | []*Scenario
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
	case *CivilianAirport:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "2" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 0.000000 {
				return "Lat"
			}
			if field == 1.000000 {
				return "Lng"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Liner:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "6" {
				return "State"
			}
			if field == "13" {
				return "Timestampstring"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "Lat"
			}
			if field == 2.000000 {
				return "Lng"
			}
			if field == 3.000000 {
				return "Heading"
			}
			if field == 4.000000 {
				return "Level"
			}
			if field == 5.000000 {
				return "Speed"
			}
			if field == 7.000000 {
				return "TargetHeading"
			}
			if field == 8.000000 {
				return "TargetLocationLat"
			}
			if field == 9.000000 {
				return "TargetLocationLng"
			}
			if field == 10.000000 {
				return "DistanceToTarget"
			}
			if field == 11.000000 {
				return "MaxRotationalSpeed"
			}
			if field == 12.000000 {
				return "VerticalSpeed"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Message:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "5" {
				return "State"
			}
			if field == "6" {
				return "Name"
			}
			if field == "10" {
				return "Timestampstring"
			}
			if field == "12" {
				return "Timestampstartstring"
			}
			if field == "13" {
				return "Source"
			}
			if field == "14" {
				return "Destination"
			}
			if field == "15" {
				return "Content"
			}
			if field == "16" {
				return "About_string"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 11 {
				return "DurationSinceSimulationStart"
			}
		case float64:
			// insertion point for field dependant name
			if field == 0.000000 {
				return "Lat"
			}
			if field == 1.000000 {
				return "Lng"
			}
			if field == 2.000000 {
				return "Heading"
			}
			if field == 3.000000 {
				return "Level"
			}
			if field == 4.000000 {
				return "Speed"
			}
			if field == 7.000000 {
				return "TargetLocationLat"
			}
			if field == 8.000000 {
				return "TargetLocationLng"
			}
			if field == 9.000000 {
				return "DistanceToTarget"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "Display"
			}
		}
	case *OpsLine:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "1" {
				return "TransmissionMessage"
			}
			if field == "3" {
				return "TransmissionMessageBackward"
			}
			if field == "5" {
				return "State"
			}
			if field == "6" {
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
			if field == false {
				return "IsTransmitting"
			}
			if field == true {
				return "IsTransmittingBackward"
			}
		}
	case *Radar:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "State"
			}
			if field == "1" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 2.000000 {
				return "Lat"
			}
			if field == 3.000000 {
				return "Lng"
			}
			if field == 4.000000 {
				return "Range"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Satellite:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Line1"
			}
			if field == "2" {
				return "Line2"
			}
			if field == "9" {
				return "Timestampstring"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 3.000000 {
				return "Lat"
			}
			if field == 4.000000 {
				return "Lng"
			}
			if field == 5.000000 {
				return "Heading"
			}
			if field == 6.000000 {
				return "Level"
			}
			if field == 7.000000 {
				return "Speed"
			}
			if field == 8.000000 {
				return "VerticalSpeed"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Scenario:
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
			if field == 1.000000 {
				return "Lat"
			}
			if field == 2.000000 {
				return "Lng"
			}
			if field == 3.000000 {
				return "ZoomLevel"
			}
			if field == 4.000000 {
				return "MessageVisualSpeed"
			}
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
