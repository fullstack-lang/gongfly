package models

import (
	"fmt"
	"time"
)

type DummyData struct {
	Name                     string
	DummyString              string
	DummyInt                 int
	DummyFloat               float64
	DummyBool                bool
	DummyTime                time.Time
	DummyDuration            time.Duration
	DummyPointerToGongStruct *AnotherDummyData
}

func GenerateTableOfDummnies() (element *Element) {
	element = (&Element{Name: "TableOfDummnies "}).Stage()
	element.Type = TABLE

	// fill up title cells
	titleRow := (&Row{Name: "Title Row "}).Stage()
	element.Rows = append(element.Rows, titleRow)

	{
		cell := (&Cell{Name: "Name "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyString "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyFloat "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyInt "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyBool "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyTime "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyDuration "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyPointerToGongStruct "}).Stage()
		titleRow.Cells = append(titleRow.Cells, cell)
	}

	for dummy := range Stage.DummyDatas {
		// fill up title cells
		row := (&Row{Name: "Date Row of " + dummy.Name}).Stage()
		element.Rows = append(element.Rows, row)

		{
			cell := (&Cell{Name: dummy.Name}).Stage()
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: dummy.DummyString}).Stage()
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: fmt.Sprintf("%f", dummy.DummyFloat)}).Stage()
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: fmt.Sprintf("%d", dummy.DummyInt)}).Stage()
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: fmt.Sprintf("%t", dummy.DummyBool)}).Stage()
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: dummy.DummyTime.String()}).Stage()
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: dummy.DummyDuration.String()}).Stage()
			row.Cells = append(row.Cells, cell)
		}
		{
			if dummy.DummyPointerToGongStruct != nil {
				cell := (&Cell{Name: dummy.DummyPointerToGongStruct.Name}).Stage()
				row.Cells = append(row.Cells, cell)
			}
		}
	}
	return
}
