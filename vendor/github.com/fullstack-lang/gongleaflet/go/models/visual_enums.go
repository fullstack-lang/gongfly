package models

// VisualColorEnum ..
// swagger:enum VisualColorEnum
type VisualColorEnum string

const (
	LIGHT_BROWN_8D6E63 VisualColorEnum = "LIGHT_BROWN_8D6E63" //
	RED                VisualColorEnum = "RED"
	GREY               VisualColorEnum = "GREY"
	GREEN              VisualColorEnum = "GREEN"
	BLUE               VisualColorEnum = "BLUE"
	NONE               VisualColorEnum = "NONE"
)

// DashStyleEnum ..
// swagger:enum DashStyleEnum
type DashStyleEnum string

const (
	FIVE_TEN    DashStyleEnum = "FIVE_TEN"
	FIVE_TWENTY DashStyleEnum = "FIVE_TWENTY"
)

// VisualLayerEnum ..
// swagger:enum VisualLayerEnum
type VisualLayerEnum string

const (
	OPS     VisualLayerEnum = "OPS"
	NETWORK VisualLayerEnum = "NETWORK"
)
