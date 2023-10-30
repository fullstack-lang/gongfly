package models

// StacksNames - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum StacksNames
type StacksNames string

// values for TableExtraNameEnum
const (
	// we have three application stacks. All have the same name
	GongflyStackName      StacksNames = "gongfly"
	GongLeafleatStackName StacksNames = "gongleaflet"
	GongsimStackName      StacksNames = "gongsim"

	// we have three probes, each with 3 stacks. Their prefix have to differ
	GongflyProbeStacksPrefix     StacksNames = "gongfly-probe"
	GongsimProbeStacksPrefix     StacksNames = "gongsim-probe"
	GongleafletProbeStacksPrefix StacksNames = "gongleaflet-probe"
)
