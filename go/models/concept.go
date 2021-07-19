package models

// Concept
// swagger:model concept
type Concept struct {
	ConceptEnum ConceptEnum
}

// ConceptEnum ..
// swagger:enum ConceptEnum
type ConceptEnum string

// available concepts that enable visual filters
const (
	Aircraft_ ConceptEnum = "Aircrafts"
	Network_  ConceptEnum = "Networks"
	Center_   ConceptEnum = "Centers"
	System_   ConceptEnum = "Systems"
)

//
type ConceptInterface interface {
	GetConcept() ConceptEnum
}
