package models

import (
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

// A Report is the event that will change the operational communication is a non working mode
// swagger:model Report
type Report struct {
	gongsim_models.Event
	ReportMessage string
	Number        int
	Type          ReportEnum
	About         *Liner   // the liner the report is about
	OpsLine       *OpsLine // the reporting line
}

// ReportEnum ..
// swagger:enum ReportEnum
type ReportEnum string

// state
const (
	TAKE_OFF_COMPLETED ReportEnum = "TAKE_OFF_COMPLETED"
)
