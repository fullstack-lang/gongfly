package models

import (
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

// A Order is the event that will change the operational communication is a non working mode
// swagger:model Order
type Order struct {
	gongsim_models.Event
	OrderMessage string
	Number       int
	Type         OrderEnum

	Target    *Liner // the liner the order is about
	TargetLat float64
	TargetLng float64
}

// OrderEnum ..
// swagger:enum OrderEnum
type OrderEnum string

// state
const (
	TAKE_OFF OrderEnum = "TAKE_OFF"
)
