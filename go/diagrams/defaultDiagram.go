package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongfly/go/models"
)

var defaultDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.CenterConcept{}),
			Position: &uml.Position{
				X: 640.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.CenterConcept{}.Lat,
				},
				{
					Field: models.CenterConcept{}.Lng,
				},
			},
		},
		{
			Struct: &(models.CivilianAirport{}),
			Position: &uml.Position{
				X: 2740.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.CivilianAirport{}.Lat,
				},
				{
					Field: models.CivilianAirport{}.Lng,
				},
				{
					Field: models.CivilianAirport{}.Name,
				},
				{
					Field: models.CivilianAirport{}.TechName,
				},
			},
		},
		{
			Struct: &(models.Concept{}),
			Position: &uml.Position{
				X: 940.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Concept{}.ConceptEnum,
				},
			},
		},
		{
			Struct: new(models.ConceptEnum),
			Position: &uml.Position{
				X: 940.000000,
				Y: 500.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Aircraft_,
				},
				{
					Field: models.Center_,
				},
				{
					Field: models.Network_,
				},
				{
					Field: models.System_,
				},
			},
		},
		{
			Struct: &(models.Liner{}),
			Position: &uml.Position{
				X: 1240.000000,
				Y: 40.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Liner{}.ReporingLine,
					Middlevertice: &uml.Vertice{
						X: 1790.000000,
						Y: 200.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Liner{}.DistanceToTarget,
				},
				{
					Field: models.Liner{}.Heading,
				},
				{
					Field: models.Liner{}.Lat,
				},
				{
					Field: models.Liner{}.Level,
				},
				{
					Field: models.Liner{}.Lng,
				},
				{
					Field: models.Liner{}.MaxRotationalSpeed,
				},
				{
					Field: models.Liner{}.Name,
				},
				{
					Field: models.Liner{}.Speed,
				},
				{
					Field: models.Liner{}.State,
				},
				{
					Field: models.Liner{}.TargetHeading,
				},
				{
					Field: models.Liner{}.TargetLocationLat,
				},
				{
					Field: models.Liner{}.TargetLocationLng,
				},
				{
					Field: models.Liner{}.TechName,
				},
				{
					Field: models.Liner{}.Timestampstring,
				},
				{
					Field: models.Liner{}.VerticalSpeed,
				},
			},
		},
		{
			Struct: new(models.LinerStateEnum),
			Position: &uml.Position{
				X: 1240.000000,
				Y: 500.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.EN_ROUTE_NOMINAL,
				},
				{
					Field: models.LANDED,
				},
			},
		},
		{
			Struct: &(models.Message{}),
			Position: &uml.Position{
				X: 3040.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Message{}.About_string,
				},
				{
					Field: models.Message{}.Content,
				},
				{
					Field: models.Message{}.Destination,
				},
				{
					Field: models.Message{}.Display,
				},
				{
					Field: models.Message{}.DistanceToTarget,
				},
				{
					Field: models.Message{}.DurationSinceSimulationStart,
				},
				{
					Field: models.Message{}.Heading,
				},
				{
					Field: models.Message{}.Lat,
				},
				{
					Field: models.Message{}.Level,
				},
				{
					Field: models.Message{}.Lng,
				},
				{
					Field: models.Message{}.Name,
				},
				{
					Field: models.Message{}.Source,
				},
				{
					Field: models.Message{}.Speed,
				},
				{
					Field: models.Message{}.State,
				},
				{
					Field: models.Message{}.TargetLocationLat,
				},
				{
					Field: models.Message{}.TargetLocationLng,
				},
				{
					Field: models.Message{}.TechName,
				},
				{
					Field: models.Message{}.Timestampstartstring,
				},
				{
					Field: models.Message{}.Timestampstring,
				},
			},
		},
		{
			Struct: new(models.MessageStateEnum),
			Position: &uml.Position{
				X: 1540.000000,
				Y: 500.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.MESSAGE_ARRIVED,
				},
				{
					Field: models.MESSAGE_EN_ROUTE,
				},
			},
		},
		{
			Struct: &(models.MovingObject{}),
			Position: &uml.Position{
				X: 3340.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.MovingObject{}.Heading,
				},
				{
					Field: models.MovingObject{}.Lat,
				},
				{
					Field: models.MovingObject{}.Level,
				},
				{
					Field: models.MovingObject{}.Lng,
				},
				{
					Field: models.MovingObject{}.Speed,
				},
			},
		},
		{
			Struct: &(models.Network{}),
			Position: &uml.Position{
				X: 1540.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Network{}.IsTransmitting,
				},
				{
					Field: models.Network{}.IsTransmittingBackward,
				},
				{
					Field: models.Network{}.TransmissionMessage,
				},
				{
					Field: models.Network{}.TransmissionMessageBackward,
				},
			},
		},
		{
			Struct: new(models.OperationalLineStateEnum),
			Position: &uml.Position{
				X: 1840.000000,
				Y: 500.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.OPS_COM_LINK_OPERATIONAL_LINE_NOT_WORKING,
				},
				{
					Field: models.OPS_COM_LINK_OPERATIONAL_LINE_WORKING,
				},
			},
		},
		{
			Struct: &(models.OpsLine{}),
			Position: &uml.Position{
				X: 3640.000000,
				Y: 40.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.OpsLine{}.Scenario,
					Middlevertice: &uml.Vertice{
						X: 4190.000000,
						Y: 200.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.OpsLine{}.IsTransmitting,
				},
				{
					Field: models.OpsLine{}.IsTransmittingBackward,
				},
				{
					Field: models.OpsLine{}.Name,
				},
				{
					Field: models.OpsLine{}.State,
				},
				{
					Field: models.OpsLine{}.TechName,
				},
				{
					Field: models.OpsLine{}.TransmissionMessage,
				},
				{
					Field: models.OpsLine{}.TransmissionMessageBackward,
				},
			},
		},
		{
			Struct: &(models.Order{}),
			Position: &uml.Position{
				X: 1840.000000,
				Y: 40.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Order{}.Target,
					Middlevertice: &uml.Vertice{
						X: 2390.000000,
						Y: 200.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Order{}.Duration,
				},
				{
					Field: models.Order{}.Name,
				},
				{
					Field: models.Order{}.Number,
				},
				{
					Field: models.Order{}.OrderMessage,
				},
				{
					Field: models.Order{}.TargetLat,
				},
				{
					Field: models.Order{}.TargetLng,
				},
				{
					Field: models.Order{}.Type,
				},
			},
		},
		{
			Struct: new(models.OrderEnum),
			Position: &uml.Position{
				X: 40.000000,
				Y: 500.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.TAKE_OFF,
				},
			},
		},
		{
			Struct: &(models.Radar{}),
			Position: &uml.Position{
				X: 2140.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Radar{}.Lat,
				},
				{
					Field: models.Radar{}.Lng,
				},
				{
					Field: models.Radar{}.Name,
				},
				{
					Field: models.Radar{}.Range,
				},
				{
					Field: models.Radar{}.State,
				},
				{
					Field: models.Radar{}.TechName,
				},
			},
		},
		{
			Struct: new(models.RadarStateEnum),
			Position: &uml.Position{
				X: 640.000000,
				Y: 500.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.WORKING,
				},
			},
		},
		{
			Struct: &(models.Report{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 40.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Report{}.About,
					Middlevertice: &uml.Vertice{
						X: 590.000000,
						Y: 200.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Report{}.OpsLine,
					Middlevertice: &uml.Vertice{
						X: 590.000000,
						Y: 250.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Report{}.Duration,
				},
				{
					Field: models.Report{}.Name,
				},
				{
					Field: models.Report{}.Number,
				},
				{
					Field: models.Report{}.ReportMessage,
				},
				{
					Field: models.Report{}.Type,
				},
			},
		},
		{
			Struct: new(models.ReportEnum),
			Position: &uml.Position{
				X: 340.000000,
				Y: 500.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.TAKE_OFF_COMPLETED,
				},
			},
		},
		{
			Struct: &(models.Scenario{}),
			Position: &uml.Position{
				X: 340.000000,
				Y: 40.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Scenario{}.Lat,
				},
				{
					Field: models.Scenario{}.Lng,
				},
				{
					Field: models.Scenario{}.MessageVisualSpeed,
				},
				{
					Field: models.Scenario{}.Name,
				},
				{
					Field: models.Scenario{}.ZoomLevel,
				},
			},
		},
		{
			Struct: &(models.System{}),
			Position: &uml.Position{
				X: 2440.000000,
				Y: 40.000000,
			},
		},
	},
}
