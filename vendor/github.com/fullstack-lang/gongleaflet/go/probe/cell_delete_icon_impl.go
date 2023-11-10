// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Circle:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DivIcon:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LayerGroup:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LayerGroupUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.MapOptions:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Marker:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.UserClick:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.VLine:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.VisualTrack:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

