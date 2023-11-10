// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
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
	case *models.CivilianAirport:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Liner:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Message:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.OpsLine:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Radar:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Satellite:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Scenario:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

