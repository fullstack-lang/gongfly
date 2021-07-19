package models

type System struct{}

func (system *System) GetConcept() ConceptEnum           { return System_ }
func (system *System) GetVisualLayerName() (name string) { return string(System_) }
