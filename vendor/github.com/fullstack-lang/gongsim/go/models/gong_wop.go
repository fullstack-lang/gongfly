// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type DummyAgent_WOP struct {
	// insertion point
	TechName string
	Name string
}

type Engine_WOP struct {
	// insertion point
	Name string
	EndTime string
	CurrentTime string
	SecondsSinceStart float64
	Fired int
	ControlMode ControlMode
	State EngineState
	Speed float64
}

type Event_WOP struct {
	// insertion point
	Name string
	Duration time.Duration
}

type GongsimCommand_WOP struct {
	// insertion point
	Name string
	Command GongsimCommandType
	CommandDate string
	SpeedCommandType SpeedCommandType
	DateSpeedCommand string
}

type GongsimStatus_WOP struct {
	// insertion point
	Name string
	CurrentCommand GongsimCommandType
	CompletionDate string
	CurrentSpeedCommand SpeedCommandType
	SpeedCommandCompletionDate string
}

