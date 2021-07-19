package icons

import (
	_ "embed"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

//go:embed airplane.svg
var airplane string
var Airplane *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "Airplane",
	SVG:  airplane,
}).StageCopy()

//go:embed Airport.svg
var airport string
var Airport *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "Airport",
	SVG:  airport,
}).StageCopy()

//go:embed antena.svg
var antena string
var Antena *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "Antena",
	SVG:  antena,
}).StageCopy()

//go:embed message.svg
var message string
var Message *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "Message",
	SVG:  message,
}).StageCopy()

//go:embed air_traffic_controler.svg
var air_traffic_controler string
var AirTrafficControler *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "AirTrafficControler",
	SVG:  air_traffic_controler,
}).StageCopy()

//go:embed dot_blur.svg
var dot_blur string
var DotBlur *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "DotBlur",
	SVG:  dot_blur,
}).StageCopy()

//go:embed radar.svg
var radar string
var Radar *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "Radar",
	SVG:  radar,
}).StageCopy()

//go:embed arrow_simple.svg
var arrow_simple string
var Arrow *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "Arrow",
	SVG:  arrow_simple,
}).StageCopy()

//go:embed cross_rot45.svg
var cross_rot45 string
var Cross *gongleaflet_models.VisualIcon = (&gongleaflet_models.VisualIcon{
	Name: "Cross",
	SVG:  cross_rot45,
}).StageCopy()
