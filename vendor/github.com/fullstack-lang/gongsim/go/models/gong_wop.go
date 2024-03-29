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

func (from *DummyAgent) CopyBasicFields(to *DummyAgent) {
	// insertion point
	to.TechName = from.TechName
	to.Name = from.Name
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

func (from *Engine) CopyBasicFields(to *Engine) {
	// insertion point
	to.Name = from.Name
	to.EndTime = from.EndTime
	to.CurrentTime = from.CurrentTime
	to.SecondsSinceStart = from.SecondsSinceStart
	to.Fired = from.Fired
	to.ControlMode = from.ControlMode
	to.State = from.State
	to.Speed = from.Speed
}

type Event_WOP struct {
	// insertion point
	Name string
	Duration time.Duration
}

func (from *Event) CopyBasicFields(to *Event) {
	// insertion point
	to.Name = from.Name
	to.Duration = from.Duration
}

type GongsimCommand_WOP struct {
	// insertion point
	Name string
	Command GongsimCommandType
	CommandDate string
	SpeedCommandType SpeedCommandType
	DateSpeedCommand string
}

func (from *GongsimCommand) CopyBasicFields(to *GongsimCommand) {
	// insertion point
	to.Name = from.Name
	to.Command = from.Command
	to.CommandDate = from.CommandDate
	to.SpeedCommandType = from.SpeedCommandType
	to.DateSpeedCommand = from.DateSpeedCommand
}

type GongsimStatus_WOP struct {
	// insertion point
	Name string
	CurrentCommand GongsimCommandType
	CompletionDate string
	CurrentSpeedCommand SpeedCommandType
	SpeedCommandCompletionDate string
}

func (from *GongsimStatus) CopyBasicFields(to *GongsimStatus) {
	// insertion point
	to.Name = from.Name
	to.CurrentCommand = from.CurrentCommand
	to.CompletionDate = from.CompletionDate
	to.CurrentSpeedCommand = from.CurrentSpeedCommand
	to.SpeedCommandCompletionDate = from.SpeedCommandCompletionDate
}

