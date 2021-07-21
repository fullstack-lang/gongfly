package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongfly/go/models"
)

var gongfly uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.CivilianAirport{}),
			Position: &uml.Position{
				X: 92.000000,
				Y: 39.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.Liner{}),
			Position: &uml.Position{
				X: 100.000000,
				Y: 140.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Liner{}.ReporingLine,
					Middlevertice: &uml.Vertice{
						X: 460.000000,
						Y: 236.500000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Liner{}.Name,
				},
			},
		},
		{
			Struct: new(models.LinerStateEnum),
			Position: &uml.Position{
				X: 450.000000,
				Y: 140.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
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
			Struct: &(models.MovingObject{}),
			Position: &uml.Position{
				X: 450.000000,
				Y: 280.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.OpsLine{}),
			Position: &uml.Position{
				X: 100.000000,
				Y: 270.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
	},
}
