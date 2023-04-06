package models

// GongsimStatus is the struct of the instance that is updated by the front for issuing Statuss
// swagger:model GongsimStatus
type GongsimStatus struct {
	Name                       string
	CurrentCommand             GongsimCommandType
	CompletionDate             string
	CurrentSpeedCommand        SpeedCommandType
	SpeedCommandCompletionDate string
}
