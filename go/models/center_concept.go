package models

type CenterConcept struct {

	// fixed position
	Lat float64
	Lng float64
}

func (center *CenterConcept) GetConcept() ConceptEnum           { return Center_ }
func (center *CenterConcept) GetVisualLayerName() (name string) { return string(Center_) }
