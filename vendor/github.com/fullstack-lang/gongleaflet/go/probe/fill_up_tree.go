package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
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
		case "Circle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Circle](probe.stageOfInterest)
			for _circle := range set {
				nodeInstance := (&tree.Node{Name: _circle.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_circle, "Circle", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DivIcon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DivIcon](probe.stageOfInterest)
			for _divicon := range set {
				nodeInstance := (&tree.Node{Name: _divicon.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_divicon, "DivIcon", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LayerGroup":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LayerGroup](probe.stageOfInterest)
			for _layergroup := range set {
				nodeInstance := (&tree.Node{Name: _layergroup.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_layergroup, "LayerGroup", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LayerGroupUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LayerGroupUse](probe.stageOfInterest)
			for _layergroupuse := range set {
				nodeInstance := (&tree.Node{Name: _layergroupuse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_layergroupuse, "LayerGroupUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MapOptions":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MapOptions](probe.stageOfInterest)
			for _mapoptions := range set {
				nodeInstance := (&tree.Node{Name: _mapoptions.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_mapoptions, "MapOptions", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Marker":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Marker](probe.stageOfInterest)
			for _marker := range set {
				nodeInstance := (&tree.Node{Name: _marker.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_marker, "Marker", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "UserClick":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.UserClick](probe.stageOfInterest)
			for _userclick := range set {
				nodeInstance := (&tree.Node{Name: _userclick.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_userclick, "UserClick", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "VLine":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.VLine](probe.stageOfInterest)
			for _vline := range set {
				nodeInstance := (&tree.Node{Name: _vline.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_vline, "VLine", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "VisualTrack":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.VisualTrack](probe.stageOfInterest)
			for _visualtrack := range set {
				nodeInstance := (&tree.Node{Name: _visualtrack.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_visualtrack, "VisualTrack", probe)

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
