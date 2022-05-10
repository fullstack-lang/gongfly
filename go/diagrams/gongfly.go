package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongfly/go/models"
)

var gongfly uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Liner{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 20.000000,
			},
			Width:  240.000000,
			Heigth: 273.000000,
			Links: []*uml.Link{
				{
					Field: models.Liner{}.ReporingLine,
					Middlevertice: &uml.Vertice{
						X: 410.000000,
						Y: 266.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
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
			Struct: &(models.OpsLine{}),
			Position: &uml.Position{
				X: 440.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 168.000000,
			Links: []*uml.Link{
				{
					Field: models.OpsLine{}.Scenario,
					Middlevertice: &uml.Vertice{
						X: 575.000000,
						Y: 509.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
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
			Struct: &(models.Scenario{}),
			Position: &uml.Position{
				X: 30.000000,
				Y: 450.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
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
	},
}
