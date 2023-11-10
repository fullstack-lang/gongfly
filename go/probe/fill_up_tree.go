package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
)

func fillUpTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "CivilianAirport":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CivilianAirport](probe.stageOfInterest)
			for _civilianairport := range set {
				nodeInstance := (&tree.Node{Name: _civilianairport.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_civilianairport, "CivilianAirport", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Liner":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Liner](probe.stageOfInterest)
			for _liner := range set {
				nodeInstance := (&tree.Node{Name: _liner.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_liner, "Liner", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Message":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Message](probe.stageOfInterest)
			for _message := range set {
				nodeInstance := (&tree.Node{Name: _message.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_message, "Message", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "OpsLine":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.OpsLine](probe.stageOfInterest)
			for _opsline := range set {
				nodeInstance := (&tree.Node{Name: _opsline.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_opsline, "OpsLine", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Radar":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Radar](probe.stageOfInterest)
			for _radar := range set {
				nodeInstance := (&tree.Node{Name: _radar.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_radar, "Radar", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Satellite":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Satellite](probe.stageOfInterest)
			for _satellite := range set {
				nodeInstance := (&tree.Node{Name: _satellite.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_satellite, "Satellite", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Scenario":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Scenario](probe.stageOfInterest)
			for _scenario := range set {
				nodeInstance := (&tree.Node{Name: _scenario.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_scenario, "Scenario", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
