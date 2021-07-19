package visuals

import (
	"log"

	"github.com/fullstack-lang/gongfly/go/icons"
	target_models "github.com/fullstack-lang/gongfly/go/models"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

// attachVisualTrack attaches a visual track to track
func attachVisualTrack(track gongleaflet_models.VisualTrackInterface,
	visualIcon *gongleaflet_models.VisualIcon,
	visualColorEnum gongleaflet_models.VisualColorEnum,
	displayTrackHistory bool,
	displayLevelAndSpeed bool) {

	// sometimes, the visual icon is nil (not reproductible bug)
	if visualIcon == nil {
		log.Fatal("nil visual icon")
	}

	visualTrack := new(gongleaflet_models.VisualTrack).Stage()
	display := true
	visualTrack.Display = display
	visualTrack.VisualTrackInterface = track
	visualTrack.VisualIcon = visualIcon
	visualTrack.DisplayTrackHistory = displayTrackHistory
	visualTrack.DisplayLevelAndSpeed = displayLevelAndSpeed
	visualTrack.VisualColorEnum = visualColorEnum
	visualTrack.UpdateTrack()
}

// attach visual center to center
func attachVisualCenter(
	visualCenterInterface gongleaflet_models.VisualCenterInterface,
	visualColorEnum gongleaflet_models.VisualColorEnum,
	visualIcon *gongleaflet_models.VisualIcon) {
	if visualIcon == nil {
		log.Fatal("nil visual icon")
	}
	visualCenter := new(gongleaflet_models.VisualCenter).Stage()
	visualCenter.VisualCenterInteface = visualCenterInterface
	visualCenter.VisualColorEnum = visualColorEnum
	visualCenter.VisualIcon = visualIcon
	visualCenter.UpdateCenter()
}

// attach visual line to line
func attachVisualLine(
	visualLineInterface gongleaflet_models.VisualLineInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {
	visualLine := new(gongleaflet_models.VisualLine).Stage()
	visualLine.DashStyleEnum = DashStyleEnum
	visualLine.VisualLineInterface = visualLineInterface
	visualLine.UpdateLine()
}

// attach visual circle to circle
func attachVisualCircle(
	visualCircleInterface gongleaflet_models.VisualCircleInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {
	visualCircle := new(gongleaflet_models.VisualCircle).Stage()
	visualCircle.DashStyleEnum = DashStyleEnum
	visualCircle.Circle = visualCircleInterface
	visualCircle.UpdateCircle()
}

func AttachVisualElementsToModelElements() {

	for obj := range target_models.Stage.Messages {
		attachVisualTrack(obj, icons.Arrow, gongleaflet_models.GREY, false, false)
	}
	for obj := range target_models.Stage.Liners {
		attachVisualTrack(obj, icons.Airplane, gongleaflet_models.GREY, true, true)
	}
	for obj := range target_models.Stage.Radars {
		attachVisualCenter(obj,
			gongleaflet_models.LIGHT_BROWN_8D6E63,
			icons.Radar)
		attachVisualCircle(obj, gongleaflet_models.FIVE_TWENTY)
	}
	for obj := range target_models.Stage.CivilianAirports {
		attachVisualCenter(obj, gongleaflet_models.BLUE, icons.Airport)
	}
	for obj := range target_models.Stage.OpsLines {
		attachVisualLine(obj, gongleaflet_models.FIVE_TWENTY)
	}

}
