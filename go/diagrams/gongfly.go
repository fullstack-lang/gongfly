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
				X: 100.000000,
				Y: 140.000000,
			},
			Width:  240.000000,
			Heigth: 273.000000,
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
			Struct: &(models.Message{}),
			Position: &uml.Position{
				X: 700.000000,
				Y: 180.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.MovingObject{}),
			Position: &uml.Position{
				X: 390.000000,
				Y: 230.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
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
			Struct: &(models.Scenario{}),
			Position: &uml.Position{
				X: 390.000000,
				Y: 60.000000,
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
