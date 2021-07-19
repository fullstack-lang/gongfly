package reference

import (
	"time"

	"github.com/fullstack-lang/gongfly/go/models"
)

var Scenario1 = (&models.Scenario{
	Name: "Scénario 1",

	ZoomLevel: 7,

	Lat: (CDG_LFPG.Lat + TLN_LFTH_ref.Lat) / 2.0,
	Lng: CDG_LFPG.Lng,

	MessageVisualSpeed: 5000,
}).Stage()

func init() {
	Scenario1.SetStart(time.Date(2020, time.January, 1, 6, 0, 0, 0, time.UTC))
	Scenario1.SetEnd(time.Date(2020, time.January, 1, 11, 0, 0, 0, time.UTC))
}
