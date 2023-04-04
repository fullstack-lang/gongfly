package visuals

import (
	"log"

	gongfly_icons "github.com/fullstack-lang/gongfly/go/icons"
	gongfly_models "github.com/fullstack-lang/gongfly/go/models"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

// attachVisualTrack attaches a visual track to track
func attachVisualTrack(
	gongleafletStage *gongleaflet_models.StageStruct,
	track gongleaflet_models.VisualTrackInterface,
	divIcon *gongleaflet_models.DivIcon,
	colorEnum gongleaflet_models.ColorEnum,
	layerGroup *gongleaflet_models.LayerGroup,
	displayTrackHistory bool,
	displayLevelAndSpeed bool) {

	// sometimes, the visual icon is nil (not reproductible bug)
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}

	visualTrack := new(gongleaflet_models.VisualTrack).Stage(gongleafletStage)
	visualTrack.VisualTrackInterface = track
	visualTrack.DivIcon = divIcon
	visualTrack.DisplayTrackHistory = displayTrackHistory
	visualTrack.DisplayLevelAndSpeed = displayLevelAndSpeed
	visualTrack.ColorEnum = colorEnum
	visualTrack.LayerGroup = layerGroup
	visualTrack.Stage_ = gongleafletStage
	visualTrack.UpdateTrack()
}

// attach visual center to center
func attachMarker(
	gongleafletStage *gongleaflet_models.StageStruct,
	visualCenterInterface gongleaflet_models.MarkerInterface,
	colorEnum gongleaflet_models.ColorEnum,
	divIcon *gongleaflet_models.DivIcon) {
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}
	visualCenter := new(gongleaflet_models.Marker).Stage(gongleafletStage)
	visualCenter.MarkerInteface = visualCenterInterface
	visualCenter.ColorEnum = colorEnum
	visualCenter.DivIcon = divIcon
	visualCenter.Stage_ = gongleafletStage
	visualCenter.UpdateMarker()
}

// attach visual line to line
func attachLine(
	gongleafletStage *gongleaflet_models.StageStruct,
	visualLineInterface gongleaflet_models.LineInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {

	visualLine := new(gongleaflet_models.VLine).Stage(gongleafletStage)
	visualLine.DashStyleEnum = DashStyleEnum
	visualLine.LineInterface = visualLineInterface
	visualLine.UpdateLine()
}

// attach visual circle to circle
func attachCircle(
	gongleafletStage *gongleaflet_models.StageStruct,
	visualCircleInterface gongleaflet_models.CircleInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {
	visualCircle := new(gongleaflet_models.Circle).Stage(gongleafletStage)
	visualCircle.DashStyleEnum = DashStyleEnum
	visualCircle.Circle = visualCircleInterface
	visualCircle.UpdateCircle()
}

// AttachVisualElementsToModelElements performs the
// link between the simulation objects and their visual
// representation on the leaflet map
func AttachVisualElementsToModelElements(
	gongflyStage *gongfly_models.StageStruct,
	gongleafletStage *gongleaflet_models.StageStruct,
	layerGroup *gongleaflet_models.LayerGroup) {

	for obj := range gongflyStage.Messages {
		attachVisualTrack(gongleafletStage,
			obj, gongfly_icons.Arrow, gongleaflet_models.GREY, layerGroup, false, false)
	}
	for obj := range gongflyStage.Liners {
		attachVisualTrack(gongleafletStage,
			obj, gongfly_icons.Airplane, gongleaflet_models.GREY, layerGroup, true, true)
	}
	for obj := range gongflyStage.Satellites {
		attachVisualTrack(gongleafletStage,
			obj, gongfly_icons.Satellite, gongleaflet_models.GREY, layerGroup, true, true)
	}
	for obj := range gongflyStage.Radars {
		attachMarker(gongleafletStage,
			obj,
			gongleaflet_models.LIGHT_BROWN_8D6E63,
			gongfly_icons.Radar)
		attachCircle(gongleafletStage,
			obj, gongleaflet_models.FIVE_TWENTY)
	}
	for obj := range gongflyStage.CivilianAirports {
		attachMarker(gongleafletStage,
			obj, gongleaflet_models.BLUE, gongfly_icons.Airport)
	}
	for obj := range gongflyStage.OpsLines {
		attachLine(gongleafletStage,
			obj, gongleaflet_models.FIVE_TWENTY)
	}

}
