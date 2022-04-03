package models

import (
	"fmt"
	"log"
	"time"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	geo "github.com/kellydunn/golang-geo"
)

// Message
// swagger:model message
type Message struct {
	MovingObject // concept

	// Agent
	// swagger:ignore
	gongsim_models.Agent

	State MessageStateEnum

	Name string

	// Control of the aircraft
	TargetLocationLat float64
	TargetLocationLng float64
	DistanceToTarget  float64

	// time stamps of the message
	timestamp       time.Time
	Timestampstring string

	// time stamps of the message at the beginning
	timestampstart               time.Time
	DurationSinceSimulationStart time.Duration
	Timestampstartstring         string

	From   LatLngAgentInterface `gorm:"-"`
	Source string               // string version of from

	To          LatLngAgentInterface `gorm:"-"`
	Destination string               // string version of to

	Content      string
	About_string string

	Display bool
}

func (*Message) GetLayerGroupName() (name string) { return string(Aircraft_) }

func (Message *Message) Register() (res *Message) {
	gongsim_models.AppendToSingloton(Message)
	res = Message
	return
}

// MessageStateEnum ..
// swagger:enum MessageStateEnum
type MessageStateEnum string

// state
const (
	MESSAGE_EN_ROUTE MessageStateEnum = "MESSAGE_EN_ROUTE"
	MESSAGE_ARRIVED  MessageStateEnum = "MESSAGE_ARRIVED"
)

// FireNextEvent fire next Event
func (message *Message) FireNextEvent() {

	event, _ := message.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		message.QueueEvent(checkStateEvent)

		//
		// Update Speed & Heading
		//
		// Change heading toward target location
		messagePoint := geo.NewPoint(message.Lat, message.Lng)
		targetLocation := geo.NewPoint(message.TargetLocationLat, message.TargetLocationLng)

		message.Heading = messagePoint.BearingTo(targetLocation)

		// Update speed
		message.DistanceToTarget = targetLocation.GreatCircleDistance(messagePoint)
		if message.DistanceToTarget < 10.0 { // 10 km.
			message.Speed = 0.0
			message.Display = false

			message.State = MESSAGE_ARRIVED
		}

		// update state vector of the message
		orig := geo.NewPoint(message.Lat, message.Lng)
		elapsedTimeInHours := checkStateEvent.Period.Hours()
		distance := message.Speed * elapsedTimeInHours // since speed is in km/h, the distance in km
		nextPoint := orig.PointAtDistanceAndBearing(
			distance,
			message.Heading)

		message.Lat = nextPoint.Lat()
		message.Lng = nextPoint.Lng()

	default:
		err := fmt.Sprintf("unkown event type %T", event)
		log.Panic(err)
	}
}

// functions to satisty the visual interface for track
func (message *Message) GetLat() float64     { return message.Lat }
func (message *Message) GetLng() float64     { return message.Lng }
func (message *Message) GetHeading() float64 { return message.Heading }

// speed is displayed in tens of nautical miles
func (message *Message) GetSpeed() float64 {
	return message.Speed / 18.52
}
func (message *Message) GetVerticalSpeed() float64 { return 0 }
func (message *Message) GetLevel() float64         { return 0 }

func (message *Message) GetDisplay() bool           { return message.Display }
func (message *Message) GetVisualLayerName() string { return "" }
