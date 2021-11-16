package visuals

import (
	"log"

	target_models "github.com/fullstack-lang/gongfly/go/models"

	gongleaflet_icons "github.com/fullstack-lang/gongleaflet/go/icons"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

// attachVisualTrack attaches a visual track to track
func attachVisualTrack(track gongleaflet_models.VisualTrackInterface,
	divIcon *gongleaflet_models.DivIcon,
	colorEnum gongleaflet_models.ColorEnum,
	layerGroup *gongleaflet_models.LayerGroup,
	displayTrackHistory bool,
	displayLevelAndSpeed bool) {

	// sometimes, the visual icon is nil (not reproductible bug)
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}

	visualTrack := new(gongleaflet_models.VisualTrack).Stage()
	visualTrack.VisualTrackInterface = track
	visualTrack.DivIcon = divIcon
	visualTrack.DisplayTrackHistory = displayTrackHistory
	visualTrack.DisplayLevelAndSpeed = displayLevelAndSpeed
	visualTrack.ColorEnum = colorEnum
	visualTrack.LayerGroup = layerGroup
	visualTrack.UpdateTrack()
}

// attach visual center to center
func attachMarker(
	visualCenterInterface gongleaflet_models.MarkerInterface,
	colorEnum gongleaflet_models.ColorEnum,
	divIcon *gongleaflet_models.DivIcon) {
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}
	visualCenter := new(gongleaflet_models.Marker).Stage()
	visualCenter.MarkerInteface = visualCenterInterface
	visualCenter.ColorEnum = colorEnum
	visualCenter.DivIcon = divIcon
	visualCenter.UpdateMarker()
}

// attach visual line to line
func attachLine(
	visualLineInterface gongleaflet_models.LineInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {
	visualLine := new(gongleaflet_models.VLine).Stage()
	visualLine.DashStyleEnum = DashStyleEnum
	visualLine.LineInterface = visualLineInterface
	visualLine.UpdateLine()
}

// attach visual circle to circle
func attachCircle(
	visualCircleInterface gongleaflet_models.CircleInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {
	visualCircle := new(gongleaflet_models.Circle).Stage()
	visualCircle.DashStyleEnum = DashStyleEnum
	visualCircle.Circle = visualCircleInterface
	visualCircle.UpdateCircle()
}

func AttachVisualElementsToModelElements(layerGroup *gongleaflet_models.LayerGroup) {

	for obj := range target_models.Stage.Messages {
		attachVisualTrack(obj, gongleaflet_icons.Arrow, gongleaflet_models.GREY, layerGroup, false, false)
	}
	for obj := range target_models.Stage.Liners {
		attachVisualTrack(obj, gongleaflet_icons.Airplane, gongleaflet_models.GREY, layerGroup, true, true)
	}
	for obj := range target_models.Stage.Satellites {
		attachVisualTrack(obj, gongleaflet_icons.Satellite, gongleaflet_models.GREY, layerGroup, true, true)
	}
	for obj := range target_models.Stage.Radars {
		attachMarker(obj,
			gongleaflet_models.LIGHT_BROWN_8D6E63,
			gongleaflet_icons.Radar)
		attachCircle(obj, gongleaflet_models.FIVE_TWENTY)
	}
	for obj := range target_models.Stage.CivilianAirports {
		attachMarker(obj, gongleaflet_models.BLUE, gongleaflet_icons.Airport)
	}
	for obj := range target_models.Stage.OpsLines {
		attachLine(obj, gongleaflet_models.FIVE_TWENTY)
	}

}
