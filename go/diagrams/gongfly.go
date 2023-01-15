package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongfly/go/models"
)

var gongfly uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			ReferencedGong: &(models.Liner{}),
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
					Field: models.Liner{}.Timestampstring,
				},
				{
					Field: models.Liner{}.VerticalSpeed,
				},
			},
		},
		{
			ReferencedGong: &(models.OpsLine{}),
			Position: &uml.Position{
				X: 440.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 153.000000,
			Links: []*uml.Link{
				{
					Field: models.OpsLine{}.Scenario,
					Middlevertice: &uml.Vertice{
						X: 565.000000,
						Y: 479.000000,
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
					Field: models.OpsLine{}.TransmissionMessage,
				},
				{
					Field: models.OpsLine{}.TransmissionMessageBackward,
				},
			},
		},
		{
			ReferencedGong: &(models.Satellite{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 310.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.Satellite{}.Line1,
				},
				{
					Field: models.Satellite{}.Line2,
				},
				{
					Field: models.Satellite{}.Name,
				},
			},
		},
		{
			ReferencedGong: &(models.Scenario{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 440.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
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
