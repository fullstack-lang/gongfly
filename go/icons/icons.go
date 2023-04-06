package icons

import (
	_ "embed"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

//go:embed airplane.svg
var airplane string
var Airplane *gongleaflet_models.DivIcon

//go:embed Airport.svg
var airport string
var Airport *gongleaflet_models.DivIcon

//go:embed antena.svg
var antena string
var Antena *gongleaflet_models.DivIcon

//go:embed message.svg
var message string
var Message *gongleaflet_models.DivIcon

//go:embed air_traffic_controler.svg
var air_traffic_controler string
var AirTrafficControler *gongleaflet_models.DivIcon

//go:embed dot_blur.svg
var dot_blur string
var DotBlur *gongleaflet_models.DivIcon

//go:embed radar.svg
var radar string
var Radar *gongleaflet_models.DivIcon

//go:embed arrow_simple.svg
var arrow_simple string
var Arrow *gongleaflet_models.DivIcon

//go:embed cross_rot45.svg
var cross_rot45 string
var Cross *gongleaflet_models.DivIcon

//go:embed satellite.svg
var satellite string
var Satellite *gongleaflet_models.DivIcon

func LoadIcons(gongleafletStage *gongleaflet_models.StageStruct) {

	Airplane = (&gongleaflet_models.DivIcon{
		Name: "Airplane",
		SVG:  airplane,
	}).Stage(gongleafletStage)

	Airport = (&gongleaflet_models.DivIcon{
		Name: "Airport",
		SVG:  airport,
	}).Stage(gongleafletStage)

	Antena = (&gongleaflet_models.DivIcon{
		Name: "Antena",
		SVG:  antena,
	}).Stage(gongleafletStage)

	Message = (&gongleaflet_models.DivIcon{
		Name: "Message",
		SVG:  message,
	}).Stage(gongleafletStage)

	AirTrafficControler = (&gongleaflet_models.DivIcon{
		Name: "AirTrafficControler",
		SVG:  air_traffic_controler,
	}).Stage(gongleafletStage)

	DotBlur = (&gongleaflet_models.DivIcon{
		Name: "DotBlur",
		SVG:  dot_blur,
	}).Stage(gongleafletStage)

	Radar = (&gongleaflet_models.DivIcon{
		Name: "Radar",
		SVG:  radar,
	}).Stage(gongleafletStage)

	Arrow = (&gongleaflet_models.DivIcon{
		Name: "Arrow",
		SVG:  arrow_simple,
	}).Stage(gongleafletStage)

	Cross = (&gongleaflet_models.DivIcon{
		Name: "Cross",
		SVG:  cross_rot45,
	}).Stage(gongleafletStage)

	Satellite = (&gongleaflet_models.DivIcon{
		Name: "Satellite",
		SVG:  satellite,
	}).Stage(gongleafletStage)
}
