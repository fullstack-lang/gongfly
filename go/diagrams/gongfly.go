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
			Heigth: 63.000000,
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
	},
}
