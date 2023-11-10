// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		fillUpTable[models.Circle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DivIcon" {
		fillUpTable[models.DivIcon](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LayerGroup" {
		fillUpTable[models.LayerGroup](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LayerGroupUse" {
		fillUpTable[models.LayerGroupUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MapOptions" {
		fillUpTable[models.MapOptions](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Marker" {
		fillUpTable[models.Marker](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "UserClick" {
		fillUpTable[models.UserClick](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "VLine" {
		fillUpTable[models.VLine](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "VisualTrack" {
		fillUpTable[models.VisualTrack](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
