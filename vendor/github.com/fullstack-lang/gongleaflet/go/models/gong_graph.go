// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *DivIcon:
		ok = stage.IsStagedDivIcon(target)

	case *LayerGroup:
		ok = stage.IsStagedLayerGroup(target)

	case *LayerGroupUse:
		ok = stage.IsStagedLayerGroupUse(target)

	case *MapOptions:
		ok = stage.IsStagedMapOptions(target)

	case *Marker:
		ok = stage.IsStagedMarker(target)

	case *UserClick:
		ok = stage.IsStagedUserClick(target)

	case *VLine:
		ok = stage.IsStagedVLine(target)

	case *VisualTrack:
		ok = stage.IsStagedVisualTrack(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedCircle(circle *Circle) (ok bool) {

		_, ok = stage.Circles[circle]
	
		return
	}

	func (stage *StageStruct) IsStagedDivIcon(divicon *DivIcon) (ok bool) {

		_, ok = stage.DivIcons[divicon]
	
		return
	}

	func (stage *StageStruct) IsStagedLayerGroup(layergroup *LayerGroup) (ok bool) {

		_, ok = stage.LayerGroups[layergroup]
	
		return
	}

	func (stage *StageStruct) IsStagedLayerGroupUse(layergroupuse *LayerGroupUse) (ok bool) {

		_, ok = stage.LayerGroupUses[layergroupuse]
	
		return
	}

	func (stage *StageStruct) IsStagedMapOptions(mapoptions *MapOptions) (ok bool) {

		_, ok = stage.MapOptionss[mapoptions]
	
		return
	}

	func (stage *StageStruct) IsStagedMarker(marker *Marker) (ok bool) {

		_, ok = stage.Markers[marker]
	
		return
	}

	func (stage *StageStruct) IsStagedUserClick(userclick *UserClick) (ok bool) {

		_, ok = stage.UserClicks[userclick]
	
		return
	}

	func (stage *StageStruct) IsStagedVLine(vline *VLine) (ok bool) {

		_, ok = stage.VLines[vline]
	
		return
	}

	func (stage *StageStruct) IsStagedVisualTrack(visualtrack *VisualTrack) (ok bool) {

		_, ok = stage.VisualTracks[visualtrack]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Circle:
		stage.StageBranchCircle(target)

	case *DivIcon:
		stage.StageBranchDivIcon(target)

	case *LayerGroup:
		stage.StageBranchLayerGroup(target)

	case *LayerGroupUse:
		stage.StageBranchLayerGroupUse(target)

	case *MapOptions:
		stage.StageBranchMapOptions(target)

	case *Marker:
		stage.StageBranchMarker(target)

	case *UserClick:
		stage.StageBranchUserClick(target)

	case *VLine:
		stage.StageBranchVLine(target)

	case *VisualTrack:
		stage.StageBranchVisualTrack(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if IsStaged(stage, circle) {
		return
	}

	circle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circle.LayerGroup != nil {
		StageBranch(stage, circle.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDivIcon(divicon *DivIcon) {

	// check if instance is already staged
	if IsStaged(stage, divicon) {
		return
	}

	divicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLayerGroup(layergroup *LayerGroup) {

	// check if instance is already staged
	if IsStaged(stage, layergroup) {
		return
	}

	layergroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLayerGroupUse(layergroupuse *LayerGroupUse) {

	// check if instance is already staged
	if IsStaged(stage, layergroupuse) {
		return
	}

	layergroupuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if layergroupuse.LayerGroup != nil {
		StageBranch(stage, layergroupuse.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMapOptions(mapoptions *MapOptions) {

	// check if instance is already staged
	if IsStaged(stage, mapoptions) {
		return
	}

	mapoptions.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layergroupuse := range mapoptions.LayerGroupUses {
		StageBranch(stage, _layergroupuse)
	}

}

func (stage *StageStruct) StageBranchMarker(marker *Marker) {

	// check if instance is already staged
	if IsStaged(stage, marker) {
		return
	}

	marker.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if marker.LayerGroup != nil {
		StageBranch(stage, marker.LayerGroup)
	}
	if marker.DivIcon != nil {
		StageBranch(stage, marker.DivIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchUserClick(userclick *UserClick) {

	// check if instance is already staged
	if IsStaged(stage, userclick) {
		return
	}

	userclick.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchVLine(vline *VLine) {

	// check if instance is already staged
	if IsStaged(stage, vline) {
		return
	}

	vline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if vline.LayerGroup != nil {
		StageBranch(stage, vline.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchVisualTrack(visualtrack *VisualTrack) {

	// check if instance is already staged
	if IsStaged(stage, visualtrack) {
		return
	}

	visualtrack.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if visualtrack.LayerGroup != nil {
		StageBranch(stage, visualtrack.LayerGroup)
	}
	if visualtrack.DivIcon != nil {
		StageBranch(stage, visualtrack.DivIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Circle:
		stage.UnstageBranchCircle(target)

	case *DivIcon:
		stage.UnstageBranchDivIcon(target)

	case *LayerGroup:
		stage.UnstageBranchLayerGroup(target)

	case *LayerGroupUse:
		stage.UnstageBranchLayerGroupUse(target)

	case *MapOptions:
		stage.UnstageBranchMapOptions(target)

	case *Marker:
		stage.UnstageBranchMarker(target)

	case *UserClick:
		stage.UnstageBranchUserClick(target)

	case *VLine:
		stage.UnstageBranchVLine(target)

	case *VisualTrack:
		stage.UnstageBranchVisualTrack(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if ! IsStaged(stage, circle) {
		return
	}

	circle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circle.LayerGroup != nil {
		UnstageBranch(stage, circle.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDivIcon(divicon *DivIcon) {

	// check if instance is already staged
	if ! IsStaged(stage, divicon) {
		return
	}

	divicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLayerGroup(layergroup *LayerGroup) {

	// check if instance is already staged
	if ! IsStaged(stage, layergroup) {
		return
	}

	layergroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLayerGroupUse(layergroupuse *LayerGroupUse) {

	// check if instance is already staged
	if ! IsStaged(stage, layergroupuse) {
		return
	}

	layergroupuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if layergroupuse.LayerGroup != nil {
		UnstageBranch(stage, layergroupuse.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMapOptions(mapoptions *MapOptions) {

	// check if instance is already staged
	if ! IsStaged(stage, mapoptions) {
		return
	}

	mapoptions.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layergroupuse := range mapoptions.LayerGroupUses {
		UnstageBranch(stage, _layergroupuse)
	}

}

func (stage *StageStruct) UnstageBranchMarker(marker *Marker) {

	// check if instance is already staged
	if ! IsStaged(stage, marker) {
		return
	}

	marker.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if marker.LayerGroup != nil {
		UnstageBranch(stage, marker.LayerGroup)
	}
	if marker.DivIcon != nil {
		UnstageBranch(stage, marker.DivIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchUserClick(userclick *UserClick) {

	// check if instance is already staged
	if ! IsStaged(stage, userclick) {
		return
	}

	userclick.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchVLine(vline *VLine) {

	// check if instance is already staged
	if ! IsStaged(stage, vline) {
		return
	}

	vline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if vline.LayerGroup != nil {
		UnstageBranch(stage, vline.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchVisualTrack(visualtrack *VisualTrack) {

	// check if instance is already staged
	if ! IsStaged(stage, visualtrack) {
		return
	}

	visualtrack.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if visualtrack.LayerGroup != nil {
		UnstageBranch(stage, visualtrack.LayerGroup)
	}
	if visualtrack.DivIcon != nil {
		UnstageBranch(stage, visualtrack.DivIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

