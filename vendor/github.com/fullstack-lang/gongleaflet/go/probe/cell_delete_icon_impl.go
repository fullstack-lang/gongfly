// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.playground = playground
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Circle:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.DivIcon:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.LayerGroup:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.LayerGroupUse:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.MapOptions:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Marker:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.UserClick:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.VLine:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.VisualTrack:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.playground.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.playground)
	fillUpTree(cellDeleteIconImpl.playground)
	cellDeleteIconImpl.playground.tableStage.Commit()
}

