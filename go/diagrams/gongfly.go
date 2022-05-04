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
				Y: 80.000000,
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
				X: 540.000000,
				Y: 110.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.MovingObject{}),
			Position: &uml.Position{
				X: 320.000000,
				Y: 220.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
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
				X: 290.000000,
				Y: 70.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
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
