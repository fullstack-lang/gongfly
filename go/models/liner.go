package models

import (
	"fmt"
	"log"
	"math"
	"time"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	geo "github.com/kellydunn/golang-geo"
)

// Liner is a moving object
// swagger:model liner
type Liner struct {
	Name string

	MovingObject // concept

	// Agent
	// swagger:ignore
	gongsim_models.Agent

	State LinerStateEnum

	// Control of the aircraft
	TargetHeading     float64
	TargetLocationLat float64
	TargetLocationLng float64
	DistanceToTarget  float64

	// Max rotational speed, in degrees/seconds
	MaxRotationalSpeed float64

	// Vertical Speed
	VerticalSpeed float64

	// time stamps of the liner
	timestamp       time.Time
	Timestampstring string

	ReporingLine *OpsLine
}

func (Liner *Liner) Register() (res *Liner) {
	gongsim_models.AppendToSingloton(Liner)
	res = Liner
	return
}

// LinerStateEnum ..
// swagger:enum LinerStateEnum
type LinerStateEnum string

// state
const (
	EN_ROUTE_NOMINAL LinerStateEnum = "EN_ROUTE_NOMINAL"
	LANDED           LinerStateEnum = "LANDED"
)

// FireNextEvent fire next Event
func (liner *Liner) FireNextEvent() {

	event, _ := liner.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		liner.QueueEvent(checkStateEvent)

		//
		// Update Speed & Heading
		//
		// Change heading toward target location
		linerPoint := geo.NewPoint(liner.Lat, liner.Lng)
		targetLocation := geo.NewPoint(liner.TargetLocationLat, liner.TargetLocationLng)

		liner.TargetHeading = linerPoint.BearingTo(targetLocation)
		liner.TargetHeading = math.Mod(liner.TargetHeading+360.0, 360.0)

		maxNumberOfDegreesPerStep := liner.MaxRotationalSpeed * checkStateEvent.Period.Seconds()
		desiredRotation := liner.TargetHeading - liner.Heading
		if math.Abs(desiredRotation) > maxNumberOfDegreesPerStep {
			desiredRotation = desiredRotation * (maxNumberOfDegreesPerStep / math.Abs(desiredRotation))
		}
		newHeading := math.Mod(liner.Heading+desiredRotation+360, 360.0)

		liner.Heading = newHeading
		// log.Printf("Heading %f, New heading %f ", liner.Heading, newHeading)

		// Update speed
		liner.DistanceToTarget = targetLocation.GreatCircleDistance(linerPoint)
		if liner.DistanceToTarget < 10.0 { // 10 km.
			liner.Speed = 0.0

			if liner.State != LANDED {
				log.Printf("Liner %s landed at lat %f lng %f at time %s ",
					liner.Name, liner.TargetLocationLat, liner.TargetLocationLng, checkStateEvent.GetFireTime())
				liner.State = LANDED
			}
		}

		// update state vector of the liner
		orig := geo.NewPoint(liner.Lat, liner.Lng)
		elapsedTimeInHours := checkStateEvent.Period.Hours()
		distance := liner.Speed * elapsedTimeInHours // since speed is in km/h, the distance in km
		nextPoint := orig.PointAtDistanceAndBearing(
			distance,
			liner.Heading)

		liner.Lat = nextPoint.Lat()
		liner.Lng = nextPoint.Lng()
		liner.timestamp = checkStateEvent.GetFireTime()
		liner.Timestampstring = liner.timestamp.String()

	case *Order:
		order := event.(*Order)

		log.Printf("liner %s receives order %s at %s", liner.Name, order.OrderMessage, order.GetFireTime())

	default:
		err := fmt.Sprintf("unkown event type %T", event)
		log.Panic(err)
	}
}

// functions to satisty the visual interface for track
func (liner *Liner) GetLat() float64     { return liner.Lat }
func (liner *Liner) GetLng() float64     { return liner.Lng }
func (liner *Liner) GetHeading() float64 { return liner.Heading }

// speed is displayed in tens of nautical miles
func (liner *Liner) GetSpeed() float64 {
	return liner.Speed / 18.52
}
func (liner *Liner) GetVerticalSpeed() float64 { return liner.VerticalSpeed }
func (liner *Liner) GetLevel() float64         { return liner.Level }

// specific
// func (liner *Liner) GetColorEnum() ColorEnum { return GREY }

func (liner *Liner) GetDisplay() bool { return true }

func (*Liner) GetLayerGroupName() (name string) { return string(Aircraft_) }
