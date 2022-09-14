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
	case CheckoutScheduler:
		fieldCoder := CheckoutScheduler{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.NbUpdatesFromFront = 1
		return (any)(fieldCoder).(Type)
	case Circle:
		fieldCoder := Circle{}
		// insertion point for field dependant code
		fieldCoder.Lat = 0.000000
		fieldCoder.Lng = 1.000000
		fieldCoder.Name = "2"
		fieldCoder.Radius = 3.000000
		fieldCoder.ColorEnum = "4"
		fieldCoder.DashStyleEnum = "5"
		return (any)(fieldCoder).(Type)
	case DivIcon:
		fieldCoder := DivIcon{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.SVG = "1"
		return (any)(fieldCoder).(Type)
	case LayerGroup:
		fieldCoder := LayerGroup{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.DisplayName = "1"
		return (any)(fieldCoder).(Type)
	case LayerGroupUse:
		fieldCoder := LayerGroupUse{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Display = false
		return (any)(fieldCoder).(Type)
	case MapOptions:
		fieldCoder := MapOptions{}
		// insertion point for field dependant code
		fieldCoder.Lat = 0.000000
		fieldCoder.Lng = 1.000000
		fieldCoder.Name = "2"
		fieldCoder.ZoomLevel = 3.000000
		fieldCoder.UrlTemplate = "4"
		fieldCoder.Attribution = "5"
		fieldCoder.MaxZoom = 6
		fieldCoder.ZoomControl = false
		fieldCoder.AttributionControl = true
		fieldCoder.ZoomSnap = 9
		return (any)(fieldCoder).(Type)
	case Marker:
		fieldCoder := Marker{}
		// insertion point for field dependant code
		fieldCoder.Lat = 0.000000
		fieldCoder.Lng = 1.000000
		fieldCoder.Name = "2"
		fieldCoder.ColorEnum = "3"
		return (any)(fieldCoder).(Type)
	case UserClick:
		fieldCoder := UserClick{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Lat = 1.000000
		fieldCoder.Lng = 2.000000
		fieldCoder.TimeOfClick = time.Date(3, time.January, 0, 0, 0, 0, 0, time.UTC)
		return (any)(fieldCoder).(Type)
	case VLine:
		fieldCoder := VLine{}
		// insertion point for field dependant code
		fieldCoder.StartLat = 0.000000
		fieldCoder.StartLng = 1.000000
		fieldCoder.EndLat = 2.000000
		fieldCoder.EndLng = 3.000000
		fieldCoder.Name = "4"
		fieldCoder.ColorEnum = "5"
		fieldCoder.DashStyleEnum = "6"
		fieldCoder.IsTransmitting = "8"
		fieldCoder.Message = "9"
		fieldCoder.IsTransmittingBackward = "10"
		fieldCoder.MessageBackward = "11"
		return (any)(fieldCoder).(Type)
	case VisualTrack:
		fieldCoder := VisualTrack{}
		// insertion point for field dependant code
		fieldCoder.Lat = 0.000000
		fieldCoder.Lng = 1.000000
		fieldCoder.Heading = 2.000000
		fieldCoder.Level = 3.000000
		fieldCoder.Speed = 4.000000
		fieldCoder.VerticalSpeed = 5.000000
		fieldCoder.Name = "6"
		fieldCoder.ColorEnum = "7"
		fieldCoder.DisplayTrackHistory = false
		fieldCoder.DisplayLevelAndSpeed = true
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *CheckoutScheduler | []*CheckoutScheduler | *Circle | []*Circle | *DivIcon | []*DivIcon | *LayerGroup | []*LayerGroup | *LayerGroupUse | []*LayerGroupUse | *MapOptions | []*MapOptions | *Marker | []*Marker | *UserClick | []*UserClick | *VLine | []*VLine | *VisualTrack | []*VisualTrack
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
	case *CheckoutScheduler:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "NbUpdatesFromFront"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Circle:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "2" {
				return "Name"
			}
			if field == "4" {
				return "ColorEnum"
			}
			if field == "5" {
				return "DashStyleEnum"
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
			if field == 3.000000 {
				return "Radius"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *DivIcon:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "SVG"
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
	case *LayerGroup:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "DisplayName"
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
	case *LayerGroupUse:
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
			if field == false {
				return "Display"
			}
		}
	case *MapOptions:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "2" {
				return "Name"
			}
			if field == "4" {
				return "UrlTemplate"
			}
			if field == "5" {
				return "Attribution"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 6 {
				return "MaxZoom"
			}
			if field == 9 {
				return "ZoomSnap"
			}
		case float64:
			// insertion point for field dependant name
			if field == 0.000000 {
				return "Lat"
			}
			if field == 1.000000 {
				return "Lng"
			}
			if field == 3.000000 {
				return "ZoomLevel"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "ZoomControl"
			}
			if field == true {
				return "AttributionControl"
			}
		}
	case *Marker:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "2" {
				return "Name"
			}
			if field == "3" {
				return "ColorEnum"
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
	case *UserClick:
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
		case time.Time:
			// insertion point for field dependant name
			if field == time.Date(3, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "TimeOfClick"
			}
		case bool:
			// insertion point for field dependant name
		}
	case *VLine:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "4" {
				return "Name"
			}
			if field == "5" {
				return "ColorEnum"
			}
			if field == "6" {
				return "DashStyleEnum"
			}
			if field == "8" {
				return "IsTransmitting"
			}
			if field == "9" {
				return "Message"
			}
			if field == "10" {
				return "IsTransmittingBackward"
			}
			if field == "11" {
				return "MessageBackward"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 0.000000 {
				return "StartLat"
			}
			if field == 1.000000 {
				return "StartLng"
			}
			if field == 2.000000 {
				return "EndLat"
			}
			if field == 3.000000 {
				return "EndLng"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *VisualTrack:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "6" {
				return "Name"
			}
			if field == "7" {
				return "ColorEnum"
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
			if field == 2.000000 {
				return "Heading"
			}
			if field == 3.000000 {
				return "Level"
			}
			if field == 4.000000 {
				return "Speed"
			}
			if field == 5.000000 {
				return "VerticalSpeed"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "DisplayTrackHistory"
			}
			if field == true {
				return "DisplayLevelAndSpeed"
			}
		}
	default:
		return ""
	}
	_ = field
	return ""
}
