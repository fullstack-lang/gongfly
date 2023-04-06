package models

import (
	"fmt"
	"log"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

const MessageSpeed = 1000.0 // in km/h

// OpsLine
//
// An OpsLine is an  agent simulating a operational communication between
// a Aircraft and a Center
//
// swagger:model OpsLine
type OpsLine struct {
	Network // for filtering sake and for handling messages

	// Agent
	// swagger:ignore
	gongsim_models.Agent

	Scenario *Scenario

	State OperationalLineStateEnum

	Name string

	// Authority
	Authority LatLngAgentInterface `gorm:"-"`

	// Subordonate
	Subordonate LatLngAgentInterface `gorm:"-"`
}

func (OpsLine *OpsLine) Register(engine *gongsim_models.Engine) (res *OpsLine) {
	gongsim_models.AppendToSingloton(engine, OpsLine)
	res = OpsLine
	return
}

// OperationalLineStateEnum ..
// swagger:enum OperationalLineStateEnum
type OperationalLineStateEnum string

// state
const (
	OPS_COM_LINK_OPERATIONAL_LINE_WORKING     OperationalLineStateEnum = "OPS_COM_LINK_OPERATIONAL_LINE_WORKING"
	OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING OperationalLineStateEnum = "OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING"
)

// FireNextEvent fire next Event
func (opsLine *OpsLine) FireNextEvent() {

	event, _ := opsLine.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		opsLine.QueueEvent(checkStateEvent)

	case *Order:
		order := event.(*Order)

		log.Printf("%s --> %s transmits order %s at %s",
			opsLine.Authority.GetName(),
			opsLine.Subordonate.GetName(),
			order.OrderMessage,
			order.GetFireTime())
		display := true
		message := (&Message{
			MovingObject: MovingObject{
				Lat:   opsLine.GetStartLat(),
				Lng:   opsLine.GetStartLng(),
				Speed: MessageSpeed, // km/h
			},
			timestamp:                    order.GetFireTime(),
			Timestampstring:              order.GetFireTime().String(),
			DurationSinceSimulationStart: order.GetFireTime().Sub(opsLine.Engine.GetStartTime()),
			timestampstart:               order.GetFireTime(),
			Timestampstartstring:         order.GetFireTime().String(),

			TargetLocationLat: opsLine.Subordonate.GetLat(),
			TargetLocationLng: opsLine.Subordonate.GetLng(),

			Name:        order.OrderMessage,
			From:        opsLine.Authority,
			Source:      opsLine.Authority.GetName(),
			To:          opsLine.Subordonate,
			Destination: opsLine.Subordonate.GetName(),
			Content:     order.OrderMessage,
			Display:     display,
		})
		if opsLine.Scenario != nil {
			message.Speed = opsLine.Scenario.MessageVisualSpeed
		}

		// Does not work yet
		// message.Stage(&Stage).Register(engine).QueueUpdateEvent(1 * time.Second)
		// _ = message

		// display the message
		opsLine.NetworkCallbackInterface = opsLine
		opsLine.TransmissionMessage = order.OrderMessage
		opsLine.IsTransmitting = true

		// generate order for subordonate
		opsLine.Subordonate.TransfertAndQueueEvent(order)
	case *Report:
		report := event.(*Report)

		log.Printf("%s --> %s transmits report %s at %s",
			opsLine.Authority.GetName(),
			opsLine.Subordonate.GetName(),
			report.ReportMessage,
			report.GetFireTime())

		display := true
		message := (&Message{
			MovingObject: MovingObject{
				Lat:   opsLine.Subordonate.GetLat(),
				Lng:   opsLine.Subordonate.GetLng(),
				Speed: MessageSpeed, // km/h
			},
			timestamp:                    report.GetFireTime(),
			Timestampstring:              report.GetFireTime().String(),
			DurationSinceSimulationStart: report.GetFireTime().Sub(opsLine.Engine.GetStartTime()),
			timestampstart:               report.GetFireTime(),
			Timestampstartstring:         report.GetFireTime().String(),

			TargetLocationLat: opsLine.Authority.GetLat(),
			TargetLocationLng: opsLine.Authority.GetLng(),
			Name:              report.ReportMessage,
			From:              opsLine.Subordonate,
			Source:            opsLine.Subordonate.GetName(),
			To:                opsLine.Authority,
			Destination:       opsLine.Authority.GetName(),
			Content:           report.ReportMessage,
			// About:                              report.About.GetName(),
			Display: display,
		})

		if opsLine.Scenario != nil {
			message.Speed = opsLine.Scenario.MessageVisualSpeed
		}
		// Does not work yet
		// message.Stage(&Stage).Register(engine).QueueUpdateEvent(1 * time.Second)
		// _ = message

		// display message
		opsLine.NetworkCallbackInterface = opsLine

		// generate report for authority
		opsLine.TransmissionMessageBackward = report.ReportMessage
		opsLine.IsTransmittingBackward = true

		// generate event for authority
		opsLine.Authority.TransfertAndQueueEvent(report)
	default:
		err := fmt.Sprintf("unkown event type %T", event)
		log.Panic(err)
	}
}

// functions to satisty the visual interface for center
func (link *OpsLine) GetStartLat() float64 { return link.Authority.GetLat() }
func (link *OpsLine) GetStartLng() float64 { return link.Authority.GetLng() }
func (link *OpsLine) GetEndLat() float64   { return link.Subordonate.GetLat() }
func (link *OpsLine) GetEndLng() float64   { return link.Subordonate.GetLng() }

func (link *OpsLine) GetIsTransmittingOverride() bool { return link.IsTransmitting }
func (link *OpsLine) GetMessageOverride() string      { return link.TransmissionMessage }
func (link *OpsLine) GetIsTransmittingBackwardOverride() bool {
	return link.IsTransmittingBackward
}
func (link *OpsLine) GetMessageBackwardOverride() string {
	return link.TransmissionMessageBackward
}
