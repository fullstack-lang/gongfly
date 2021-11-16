package models

import (
	"fmt"
	"log"
	"time"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	gosatellite "github.com/joshuaferrara/go-satellite"
)

// Satellite is a moving object
// swagger:model satellite
type Satellite struct {
	MovingObject // concept

	Line1 string
	Line2 string

	// Agent
	// swagger:ignore
	gongsim_models.Agent

	Name string

	// the reference simulation
	gosatellite gosatellite.Satellite

	// Control of the satelitte
	TargetHeading     float64
	TargetLocationLat float64
	TargetLocationLng float64
	DistanceToTarget  float64

	// Max rotational speed, in degrees/seconds
	MaxRotationalSpeed float64

	// Vertical Speed
	VerticalSpeed float64

	// time stamps of the satellite
	timestamp       time.Time
	Timestampstring string
}

func (satellite *Satellite) Register() (res *Satellite) {
	gongsim_models.AppendToSingloton(satellite)
	res = satellite
	return
}

func (satellite *Satellite) InitFromTLE() (res *Satellite) {
	res = satellite
	res.gosatellite = gosatellite.TLEToSat(res.Line1, res.Line2, "wgs72")
	return
}

// FireNextEvent fire next Event
func (satellite *Satellite) FireNextEvent() {

	event, _ := satellite.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		satellite.QueueEvent(checkStateEvent)

		// get the time
		// func Propagate(sat Satellite, year int, month int, day, hours, minutes, seconds int) (position, velocity Vector3) {
		julianDay := gosatellite.JDay(checkStateEvent.GetFireTime().Year(),
			int(checkStateEvent.GetFireTime().Month()),
			checkStateEvent.GetFireTime().Day(),
			checkStateEvent.GetFireTime().Hour(),
			checkStateEvent.GetFireTime().Minute(),
			checkStateEvent.GetFireTime().Second())

		thetaG := gosatellite.ThetaG_JD(julianDay)

		position, velocity :=
			gosatellite.Propagate(satellite.gosatellite,
				checkStateEvent.GetFireTime().Year(),
				int(checkStateEvent.GetFireTime().Month()),
				checkStateEvent.GetFireTime().Day(),
				checkStateEvent.GetFireTime().Hour(),
				checkStateEvent.GetFireTime().Minute(),
				checkStateEvent.GetFireTime().Second())

		// compute position in lat lng
		computedAltitude, b, latLngRad := gosatellite.ECIToLLA(position, thetaG)
		_ = computedAltitude
		_ = b
		_ = position
		_ = velocity

		// convert
		satellite.Level = computedAltitude * 3280.84

		latLngDeg := gosatellite.LatLongDeg(latLngRad)
		_ = latLngDeg

		satellite.Lat = latLngDeg.Latitude
		satellite.Lng = latLngDeg.Longitude

		satellite.timestamp = checkStateEvent.GetFireTime()
		satellite.Timestampstring = satellite.timestamp.String()

	case *Order:
		order := event.(*Order)

		log.Printf("satellite %s receives order %s at %s", satellite.Name, order.OrderMessage, order.GetFireTime())

	default:
		err := fmt.Sprintf("unkown event type %T", event)
		log.Panic(err)
	}
}

// functions to satisty the visual interface for track
func (satellite *Satellite) GetLat() float64     { return satellite.Lat }
func (satellite *Satellite) GetLng() float64     { return satellite.Lng }
func (satellite *Satellite) GetHeading() float64 { return satellite.Heading }

// speed is displayed in tens of nautical miles
func (satellite *Satellite) GetSpeed() float64 {
	return satellite.Speed / 18.52
}
func (satellite *Satellite) GetVerticalSpeed() float64 { return satellite.VerticalSpeed }
func (satellite *Satellite) GetLevel() float64         { return satellite.Level }
func (satellite *Satellite) GetName() (name string)    { return satellite.Name }

// specific
// func (satellite *Satellite) GetColorEnum() ColorEnum { return GREY }

func (satellite *Satellite) GetDisplay() bool { return true }

func (*Satellite) GetLayerGroupName() (name string) { return string(Satellite_) }
