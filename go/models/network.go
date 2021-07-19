package models

// Network is a two-ways message transmitter
type Network struct {
	IsTransmitting      bool   // true if it has to display the message
	TransmissionMessage string // message to display

	IsTransmittingBackward      bool   // true if it has to display the message
	TransmissionMessageBackward string // message to display

	// this interface can be set to override the default behavior of the network
	// which is doing nothing
	NetworkCallbackInterface NetworkCallbackInterface `gorm:"-"`
}

type NetworkCallbackInterface interface {
	GetIsTransmittingOverride() bool
	GetMessageOverride() string

	GetIsTransmittingBackwardOverride() bool
	GetMessageBackwardOverride() string
}

func (network *Network) GetConcept() ConceptEnum           { return Network_ }
func (network *Network) GetVisualLayerName() (name string) { return string(Network_) }

func (network *Network) GetIsTransmitting() bool {
	if network.NetworkCallbackInterface != nil {
		return network.NetworkCallbackInterface.GetIsTransmittingOverride()
	}
	return false
}
func (network *Network) GetMessage() (name string) {
	if network.NetworkCallbackInterface != nil {
		return network.NetworkCallbackInterface.GetMessageOverride()
	}
	return ""
}

func (network *Network) GetIsTransmittingBackward() bool {
	if network.NetworkCallbackInterface != nil {
		return network.NetworkCallbackInterface.GetIsTransmittingBackwardOverride()
	}
	return false
}
func (network *Network) GetMessageBackward() (name string) {
	if network.NetworkCallbackInterface != nil {
		return network.NetworkCallbackInterface.GetMessageBackwardOverride()
	}
	return ""
}
