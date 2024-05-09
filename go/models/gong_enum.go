// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for ConceptEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (conceptenum ConceptEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch conceptenum {
	// insertion code per enum code
	case Aircraft_:
		res = "Aircrafts"
	case Satellite_:
		res = "Satellites"
	case Network_:
		res = "Networks"
	case Center_:
		res = "Centers"
	case System_:
		res = "Systems"
	}
	return
}

func (conceptenum *ConceptEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Aircrafts":
		*conceptenum = Aircraft_
	case "Satellites":
		*conceptenum = Satellite_
	case "Networks":
		*conceptenum = Network_
	case "Centers":
		*conceptenum = Center_
	case "Systems":
		*conceptenum = System_
	default:
		return errUnkownEnum
	}
	return
}

func (conceptenum *ConceptEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Aircraft_":
		*conceptenum = Aircraft_
	case "Satellite_":
		*conceptenum = Satellite_
	case "Network_":
		*conceptenum = Network_
	case "Center_":
		*conceptenum = Center_
	case "System_":
		*conceptenum = System_
	default:
		return errUnkownEnum
	}
	return
}

func (conceptenum *ConceptEnum) ToCodeString() (res string) {

	switch *conceptenum {
	// insertion code per enum code
	case Aircraft_:
		res = "Aircraft_"
	case Satellite_:
		res = "Satellite_"
	case Network_:
		res = "Network_"
	case Center_:
		res = "Center_"
	case System_:
		res = "System_"
	}
	return
}

func (conceptenum ConceptEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Aircraft_")
	res = append(res, "Satellite_")
	res = append(res, "Network_")
	res = append(res, "Center_")
	res = append(res, "System_")

	return
}

func (conceptenum ConceptEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Aircrafts")
	res = append(res, "Satellites")
	res = append(res, "Networks")
	res = append(res, "Centers")
	res = append(res, "Systems")

	return
}

// Utility function for LinerStateEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (linerstateenum LinerStateEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch linerstateenum {
	// insertion code per enum code
	case EN_ROUTE_NOMINAL:
		res = "EN_ROUTE_NOMINAL"
	case LANDED:
		res = "LANDED"
	}
	return
}

func (linerstateenum *LinerStateEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EN_ROUTE_NOMINAL":
		*linerstateenum = EN_ROUTE_NOMINAL
	case "LANDED":
		*linerstateenum = LANDED
	default:
		return errUnkownEnum
	}
	return
}

func (linerstateenum *LinerStateEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EN_ROUTE_NOMINAL":
		*linerstateenum = EN_ROUTE_NOMINAL
	case "LANDED":
		*linerstateenum = LANDED
	default:
		return errUnkownEnum
	}
	return
}

func (linerstateenum *LinerStateEnum) ToCodeString() (res string) {

	switch *linerstateenum {
	// insertion code per enum code
	case EN_ROUTE_NOMINAL:
		res = "EN_ROUTE_NOMINAL"
	case LANDED:
		res = "LANDED"
	}
	return
}

func (linerstateenum LinerStateEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EN_ROUTE_NOMINAL")
	res = append(res, "LANDED")

	return
}

func (linerstateenum LinerStateEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EN_ROUTE_NOMINAL")
	res = append(res, "LANDED")

	return
}

// Utility function for MessageStateEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (messagestateenum MessageStateEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch messagestateenum {
	// insertion code per enum code
	case MESSAGE_EN_ROUTE:
		res = "MESSAGE_EN_ROUTE"
	case MESSAGE_ARRIVED:
		res = "MESSAGE_ARRIVED"
	}
	return
}

func (messagestateenum *MessageStateEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MESSAGE_EN_ROUTE":
		*messagestateenum = MESSAGE_EN_ROUTE
	case "MESSAGE_ARRIVED":
		*messagestateenum = MESSAGE_ARRIVED
	default:
		return errUnkownEnum
	}
	return
}

func (messagestateenum *MessageStateEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MESSAGE_EN_ROUTE":
		*messagestateenum = MESSAGE_EN_ROUTE
	case "MESSAGE_ARRIVED":
		*messagestateenum = MESSAGE_ARRIVED
	default:
		return errUnkownEnum
	}
	return
}

func (messagestateenum *MessageStateEnum) ToCodeString() (res string) {

	switch *messagestateenum {
	// insertion code per enum code
	case MESSAGE_EN_ROUTE:
		res = "MESSAGE_EN_ROUTE"
	case MESSAGE_ARRIVED:
		res = "MESSAGE_ARRIVED"
	}
	return
}

func (messagestateenum MessageStateEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MESSAGE_EN_ROUTE")
	res = append(res, "MESSAGE_ARRIVED")

	return
}

func (messagestateenum MessageStateEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MESSAGE_EN_ROUTE")
	res = append(res, "MESSAGE_ARRIVED")

	return
}

// Utility function for OperationalLineStateEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (operationallinestateenum OperationalLineStateEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch operationallinestateenum {
	// insertion code per enum code
	case OPS_COM_LINK_OPERATIONAL_LINE_WORKING:
		res = "OPS_COM_LINK_OPERATIONAL_LINE_WORKING"
	case OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING:
		res = "OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING"
	}
	return
}

func (operationallinestateenum *OperationalLineStateEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OPS_COM_LINK_OPERATIONAL_LINE_WORKING":
		*operationallinestateenum = OPS_COM_LINK_OPERATIONAL_LINE_WORKING
	case "OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING":
		*operationallinestateenum = OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING
	default:
		return errUnkownEnum
	}
	return
}

func (operationallinestateenum *OperationalLineStateEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OPS_COM_LINK_OPERATIONAL_LINE_WORKING":
		*operationallinestateenum = OPS_COM_LINK_OPERATIONAL_LINE_WORKING
	case "OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING":
		*operationallinestateenum = OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING
	default:
		return errUnkownEnum
	}
	return
}

func (operationallinestateenum *OperationalLineStateEnum) ToCodeString() (res string) {

	switch *operationallinestateenum {
	// insertion code per enum code
	case OPS_COM_LINK_OPERATIONAL_LINE_WORKING:
		res = "OPS_COM_LINK_OPERATIONAL_LINE_WORKING"
	case OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING:
		res = "OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING"
	}
	return
}

func (operationallinestateenum OperationalLineStateEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OPS_COM_LINK_OPERATIONAL_LINE_WORKING")
	res = append(res, "OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING")

	return
}

func (operationallinestateenum OperationalLineStateEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OPS_COM_LINK_OPERATIONAL_LINE_WORKING")
	res = append(res, "OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING")

	return
}

// Utility function for OrderEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (orderenum OrderEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch orderenum {
	// insertion code per enum code
	case TAKE_OFF:
		res = "TAKE_OFF"
	}
	return
}

func (orderenum *OrderEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TAKE_OFF":
		*orderenum = TAKE_OFF
	default:
		return errUnkownEnum
	}
	return
}

func (orderenum *OrderEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TAKE_OFF":
		*orderenum = TAKE_OFF
	default:
		return errUnkownEnum
	}
	return
}

func (orderenum *OrderEnum) ToCodeString() (res string) {

	switch *orderenum {
	// insertion code per enum code
	case TAKE_OFF:
		res = "TAKE_OFF"
	}
	return
}

func (orderenum OrderEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TAKE_OFF")

	return
}

func (orderenum OrderEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TAKE_OFF")

	return
}

// Utility function for RadarStateEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (radarstateenum RadarStateEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch radarstateenum {
	// insertion code per enum code
	case WORKING:
		res = "WORKING"
	}
	return
}

func (radarstateenum *RadarStateEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "WORKING":
		*radarstateenum = WORKING
	default:
		return errUnkownEnum
	}
	return
}

func (radarstateenum *RadarStateEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "WORKING":
		*radarstateenum = WORKING
	default:
		return errUnkownEnum
	}
	return
}

func (radarstateenum *RadarStateEnum) ToCodeString() (res string) {

	switch *radarstateenum {
	// insertion code per enum code
	case WORKING:
		res = "WORKING"
	}
	return
}

func (radarstateenum RadarStateEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "WORKING")

	return
}

func (radarstateenum RadarStateEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "WORKING")

	return
}

// Utility function for ReportEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (reportenum ReportEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch reportenum {
	// insertion code per enum code
	case TAKE_OFF_COMPLETED:
		res = "TAKE_OFF_COMPLETED"
	}
	return
}

func (reportenum *ReportEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TAKE_OFF_COMPLETED":
		*reportenum = TAKE_OFF_COMPLETED
	default:
		return errUnkownEnum
	}
	return
}

func (reportenum *ReportEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TAKE_OFF_COMPLETED":
		*reportenum = TAKE_OFF_COMPLETED
	default:
		return errUnkownEnum
	}
	return
}

func (reportenum *ReportEnum) ToCodeString() (res string) {

	switch *reportenum {
	// insertion code per enum code
	case TAKE_OFF_COMPLETED:
		res = "TAKE_OFF_COMPLETED"
	}
	return
}

func (reportenum ReportEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TAKE_OFF_COMPLETED")

	return
}

func (reportenum ReportEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TAKE_OFF_COMPLETED")

	return
}

// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case GongflyStackName:
		res = "gongfly"
	case GongLeafleatStackName:
		res = "gongleaflet"
	case GongsimStackName:
		res = "gongsim"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "gongfly":
		*stacksnames = GongflyStackName
	case "gongleaflet":
		*stacksnames = GongLeafleatStackName
	case "gongsim":
		*stacksnames = GongsimStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "GongflyStackName":
		*stacksnames = GongflyStackName
	case "GongLeafleatStackName":
		*stacksnames = GongLeafleatStackName
	case "GongsimStackName":
		*stacksnames = GongsimStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case GongflyStackName:
		res = "GongflyStackName"
	case GongLeafleatStackName:
		res = "GongLeafleatStackName"
	case GongsimStackName:
		res = "GongsimStackName"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "GongflyStackName")
	res = append(res, "GongLeafleatStackName")
	res = append(res, "GongsimStackName")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "gongfly")
	res = append(res, "gongleaflet")
	res = append(res, "gongsim")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | ConceptEnum | LinerStateEnum | MessageStateEnum | OperationalLineStateEnum | OrderEnum | RadarStateEnum | ReportEnum | StacksNames
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*ConceptEnum | *LinerStateEnum | *MessageStateEnum | *OperationalLineStateEnum | *OrderEnum | *RadarStateEnum | *ReportEnum | *StacksNames
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
