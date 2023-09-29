// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for ColorEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (colorenum ColorEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch colorenum {
	// insertion code per enum code
	case LIGHT_BROWN_8D6E63:
		res = "LIGHT_BROWN_8D6E63"
	case RED:
		res = "RED"
	case GREY:
		res = "GREY"
	case GREEN:
		res = "GREEN"
	case BLUE:
		res = "BLUE"
	case NONE:
		res = "NONE"
	}
	return
}

func (colorenum *ColorEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LIGHT_BROWN_8D6E63":
		*colorenum = LIGHT_BROWN_8D6E63
	case "RED":
		*colorenum = RED
	case "GREY":
		*colorenum = GREY
	case "GREEN":
		*colorenum = GREEN
	case "BLUE":
		*colorenum = BLUE
	case "NONE":
		*colorenum = NONE
	default:
		return errUnkownEnum
	}
	return
}

func (colorenum *ColorEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LIGHT_BROWN_8D6E63":
		*colorenum = LIGHT_BROWN_8D6E63
	case "RED":
		*colorenum = RED
	case "GREY":
		*colorenum = GREY
	case "GREEN":
		*colorenum = GREEN
	case "BLUE":
		*colorenum = BLUE
	case "NONE":
		*colorenum = NONE
	default:
		return errUnkownEnum
	}
	return
}

func (colorenum *ColorEnum) ToCodeString() (res string) {

	switch *colorenum {
	// insertion code per enum code
	case LIGHT_BROWN_8D6E63:
		res = "LIGHT_BROWN_8D6E63"
	case RED:
		res = "RED"
	case GREY:
		res = "GREY"
	case GREEN:
		res = "GREEN"
	case BLUE:
		res = "BLUE"
	case NONE:
		res = "NONE"
	}
	return
}

func (colorenum ColorEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LIGHT_BROWN_8D6E63")
	res = append(res, "RED")
	res = append(res, "GREY")
	res = append(res, "GREEN")
	res = append(res, "BLUE")
	res = append(res, "NONE")

	return
}

func (colorenum ColorEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LIGHT_BROWN_8D6E63")
	res = append(res, "RED")
	res = append(res, "GREY")
	res = append(res, "GREEN")
	res = append(res, "BLUE")
	res = append(res, "NONE")

	return
}

// Utility function for DashStyleEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (dashstyleenum DashStyleEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch dashstyleenum {
	// insertion code per enum code
	case FIVE_TEN:
		res = "FIVE_TEN"
	case FIVE_TWENTY:
		res = "FIVE_TWENTY"
	}
	return
}

func (dashstyleenum *DashStyleEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FIVE_TEN":
		*dashstyleenum = FIVE_TEN
	case "FIVE_TWENTY":
		*dashstyleenum = FIVE_TWENTY
	default:
		return errUnkownEnum
	}
	return
}

func (dashstyleenum *DashStyleEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FIVE_TEN":
		*dashstyleenum = FIVE_TEN
	case "FIVE_TWENTY":
		*dashstyleenum = FIVE_TWENTY
	default:
		return errUnkownEnum
	}
	return
}

func (dashstyleenum *DashStyleEnum) ToCodeString() (res string) {

	switch *dashstyleenum {
	// insertion code per enum code
	case FIVE_TEN:
		res = "FIVE_TEN"
	case FIVE_TWENTY:
		res = "FIVE_TWENTY"
	}
	return
}

func (dashstyleenum DashStyleEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FIVE_TEN")
	res = append(res, "FIVE_TWENTY")

	return
}

func (dashstyleenum DashStyleEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FIVE_TEN")
	res = append(res, "FIVE_TWENTY")

	return
}

// Utility function for Start_To_End_Enum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (start_to_end_enum Start_To_End_Enum) ToString() (res string) {

	// migration of former implementation of enum
	switch start_to_end_enum {
	// insertion code per enum code
	case FORWARD_START_TO_END:
		res = "FORWARD_START_TO_END"
	case BACKWARD_END_TO_START:
		res = "BACKWARD_START_TO_END"
	}
	return
}

func (start_to_end_enum *Start_To_End_Enum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FORWARD_START_TO_END":
		*start_to_end_enum = FORWARD_START_TO_END
	case "BACKWARD_START_TO_END":
		*start_to_end_enum = BACKWARD_END_TO_START
	default:
		return errUnkownEnum
	}
	return
}

func (start_to_end_enum *Start_To_End_Enum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FORWARD_START_TO_END":
		*start_to_end_enum = FORWARD_START_TO_END
	case "BACKWARD_END_TO_START":
		*start_to_end_enum = BACKWARD_END_TO_START
	default:
		return errUnkownEnum
	}
	return
}

func (start_to_end_enum *Start_To_End_Enum) ToCodeString() (res string) {

	switch *start_to_end_enum {
	// insertion code per enum code
	case FORWARD_START_TO_END:
		res = "FORWARD_START_TO_END"
	case BACKWARD_END_TO_START:
		res = "BACKWARD_END_TO_START"
	}
	return
}

func (start_to_end_enum Start_To_End_Enum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FORWARD_START_TO_END")
	res = append(res, "BACKWARD_END_TO_START")

	return
}

func (start_to_end_enum Start_To_End_Enum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FORWARD_START_TO_END")
	res = append(res, "BACKWARD_START_TO_END")

	return
}

// Utility function for TransmittingEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (transmittingenum TransmittingEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch transmittingenum {
	// insertion code per enum code
	case IS_TRANSMITTING:
		res = "IS_TRANSMITTING"
	case IS_NOT_TRANSMITTING:
		res = "IS_NOT_TRANSMITTING"
	}
	return
}

func (transmittingenum *TransmittingEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IS_TRANSMITTING":
		*transmittingenum = IS_TRANSMITTING
	case "IS_NOT_TRANSMITTING":
		*transmittingenum = IS_NOT_TRANSMITTING
	default:
		return errUnkownEnum
	}
	return
}

func (transmittingenum *TransmittingEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IS_TRANSMITTING":
		*transmittingenum = IS_TRANSMITTING
	case "IS_NOT_TRANSMITTING":
		*transmittingenum = IS_NOT_TRANSMITTING
	default:
		return errUnkownEnum
	}
	return
}

func (transmittingenum *TransmittingEnum) ToCodeString() (res string) {

	switch *transmittingenum {
	// insertion code per enum code
	case IS_TRANSMITTING:
		res = "IS_TRANSMITTING"
	case IS_NOT_TRANSMITTING:
		res = "IS_NOT_TRANSMITTING"
	}
	return
}

func (transmittingenum TransmittingEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IS_TRANSMITTING")
	res = append(res, "IS_NOT_TRANSMITTING")

	return
}

func (transmittingenum TransmittingEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IS_TRANSMITTING")
	res = append(res, "IS_NOT_TRANSMITTING")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | ColorEnum | DashStyleEnum | Start_To_End_Enum | TransmittingEnum
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*ColorEnum | *DashStyleEnum | *Start_To_End_Enum | *TransmittingEnum
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
