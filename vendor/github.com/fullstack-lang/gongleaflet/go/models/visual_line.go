package models

// VisualLine provides all necessary elements to the front to display a track
//
// swagger:model visualcenter
type VisualLine struct {
	StartLat, StartLng float64
	EndLat, EndLng     float64
	Name               string

	VisualColorEnum VisualColorEnum
	DashStyleEnum   DashStyleEnum

	// VisualLayer the object belongs to
	VisualLayer *VisualLayer

	IsTransmitting TransmittingEnum // display the message displacement
	Message        string           // message to display

	IsTransmittingBackward TransmittingEnum // display the message displacement on the return path
	MessageBackward        string           // message to display

	// swagger:ignore
	// access to the models instance that contains the original information
	VisualLineInterface VisualLineInterface `gorm:"-"`
}

// for the moment, the angular front end does not get booleans, therefore, we translate the
// booleans into enums

// TransmittingEnum ..
// swagger:enum TransmittingEnum
type TransmittingEnum string

// state
const (
	IS_TRANSMITTING     TransmittingEnum = "IS_TRANSMITTING"
	IS_NOT_TRANSMITTING TransmittingEnum = "IS_NOT_TRANSMITTING"
)

// Start_To_End_Enum ..
// swagger:enum Start_To_End_Enum
type Start_To_End_Enum string

// state
const (
	FORWARD_START_TO_END  Start_To_End_Enum = "FORWARD_START_TO_END"
	BACKWARD_END_TO_START Start_To_End_Enum = "BACKWARD_START_TO_END"
)

type VisualLineInterface interface {
	GetStartLat() (lat float64)
	GetStartLng() (lng float64)

	GetEndLat() (lat float64)
	GetEndLng() (lng float64)

	GetName() (name string)

	GetVisualLayerName() string

	// transmission
	GetIsTransmitting() bool // display the message displacement
	GetMessage() string      // message to display

	GetIsTransmittingBackward() bool // display the message displacement
	GetMessageBackward() string      // message to display

}

func (visualLine *VisualLine) UpdateLine() {
	if visualLine.VisualLineInterface != nil {
		visualLine.Name = visualLine.VisualLineInterface.GetName()

		visualLine.StartLat = visualLine.VisualLineInterface.GetStartLat()
		visualLine.StartLng = visualLine.VisualLineInterface.GetStartLng()

		visualLine.EndLat = visualLine.VisualLineInterface.GetEndLat()
		visualLine.EndLng = visualLine.VisualLineInterface.GetEndLng()

		visualLine.VisualLayer =
			computeVisualLayerFromVisualLayerName(visualLine.VisualLineInterface.GetVisualLayerName())

		// transmission status
		if visualLine.VisualLineInterface.GetIsTransmitting() {
			visualLine.IsTransmitting = IS_TRANSMITTING
		} else {
			visualLine.IsTransmitting = IS_NOT_TRANSMITTING
		}
		visualLine.Message = visualLine.VisualLineInterface.GetMessage()

		// transmission status
		if visualLine.VisualLineInterface.GetIsTransmittingBackward() {
			visualLine.IsTransmittingBackward = IS_TRANSMITTING
		} else {
			visualLine.IsTransmittingBackward = IS_NOT_TRANSMITTING
		}
		visualLine.MessageBackward = visualLine.VisualLineInterface.GetMessageBackward()

	}
}
