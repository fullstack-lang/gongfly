package models

import (
	"fmt"
	"log"
	"math"
	"time"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	geo "github.com/kellydunn/golang-geo"

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

		//
		// Update Speed & Heading
		//
		// Change heading toward target location
		satellitePoint := geo.NewPoint(satellite.Lat, satellite.Lng)
		targetLocation := geo.NewPoint(satellite.TargetLocationLat, satellite.TargetLocationLng)

		satellite.TargetHeading = satellitePoint.BearingTo(targetLocation)
		satellite.TargetHeading = math.Mod(satellite.TargetHeading+360.0, 360.0)

		maxNumberOfDegreesPerStep := satellite.MaxRotationalSpeed * checkStateEvent.Period.Seconds()
		desiredRotation := satellite.TargetHeading - satellite.Heading
		if math.Abs(desiredRotation) > maxNumberOfDegreesPerStep {
			desiredRotation = desiredRotation * (maxNumberOfDegreesPerStep / math.Abs(desiredRotation))
		}
		newHeading := math.Mod(satellite.Heading+desiredRotation+360, 360.0)

		satellite.Heading = newHeading
		// log.Printf("Heading %f, New heading %f ", satellite.Heading, newHeading)

		// update state vector of the satellite
		orig := geo.NewPoint(satellite.Lat, satellite.Lng)
		elapsedTimeInHours := checkStateEvent.Period.Hours()
		distance := satellite.Speed * elapsedTimeInHours // since speed is in km/h, the distance in km
		nextPoint := orig.PointAtDistanceAndBearing(
			distance,
			satellite.Heading)

		satellite.Lat = nextPoint.Lat()
		satellite.Lng = nextPoint.Lng()
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
