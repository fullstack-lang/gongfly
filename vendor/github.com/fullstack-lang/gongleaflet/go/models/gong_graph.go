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
// referenced by pointers or slices of pointers of the instance
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


// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Circle:
		toT := CopyBranchCircle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DivIcon:
		toT := CopyBranchDivIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LayerGroup:
		toT := CopyBranchLayerGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LayerGroupUse:
		toT := CopyBranchLayerGroupUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MapOptions:
		toT := CopyBranchMapOptions(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Marker:
		toT := CopyBranchMarker(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *UserClick:
		toT := CopyBranchUserClick(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VLine:
		toT := CopyBranchVLine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VisualTrack:
		toT := CopyBranchVisualTrack(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCircle(mapOrigCopy map[any]any, circleFrom *Circle) (circleTo  *Circle){

	// circleFrom has already been copied
	if _circleTo, ok := mapOrigCopy[circleFrom]; ok {
		circleTo = _circleTo.(*Circle)
		return
	}

	circleTo = new(Circle)
	mapOrigCopy[circleFrom] = circleTo
	circleFrom.CopyBasicFields(circleTo)

	//insertion point for the staging of instances referenced by pointers
	if circleFrom.LayerGroup != nil {
		circleTo.LayerGroup = CopyBranchLayerGroup(mapOrigCopy, circleFrom.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDivIcon(mapOrigCopy map[any]any, diviconFrom *DivIcon) (diviconTo  *DivIcon){

	// diviconFrom has already been copied
	if _diviconTo, ok := mapOrigCopy[diviconFrom]; ok {
		diviconTo = _diviconTo.(*DivIcon)
		return
	}

	diviconTo = new(DivIcon)
	mapOrigCopy[diviconFrom] = diviconTo
	diviconFrom.CopyBasicFields(diviconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLayerGroup(mapOrigCopy map[any]any, layergroupFrom *LayerGroup) (layergroupTo  *LayerGroup){

	// layergroupFrom has already been copied
	if _layergroupTo, ok := mapOrigCopy[layergroupFrom]; ok {
		layergroupTo = _layergroupTo.(*LayerGroup)
		return
	}

	layergroupTo = new(LayerGroup)
	mapOrigCopy[layergroupFrom] = layergroupTo
	layergroupFrom.CopyBasicFields(layergroupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLayerGroupUse(mapOrigCopy map[any]any, layergroupuseFrom *LayerGroupUse) (layergroupuseTo  *LayerGroupUse){

	// layergroupuseFrom has already been copied
	if _layergroupuseTo, ok := mapOrigCopy[layergroupuseFrom]; ok {
		layergroupuseTo = _layergroupuseTo.(*LayerGroupUse)
		return
	}

	layergroupuseTo = new(LayerGroupUse)
	mapOrigCopy[layergroupuseFrom] = layergroupuseTo
	layergroupuseFrom.CopyBasicFields(layergroupuseTo)

	//insertion point for the staging of instances referenced by pointers
	if layergroupuseFrom.LayerGroup != nil {
		layergroupuseTo.LayerGroup = CopyBranchLayerGroup(mapOrigCopy, layergroupuseFrom.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMapOptions(mapOrigCopy map[any]any, mapoptionsFrom *MapOptions) (mapoptionsTo  *MapOptions){

	// mapoptionsFrom has already been copied
	if _mapoptionsTo, ok := mapOrigCopy[mapoptionsFrom]; ok {
		mapoptionsTo = _mapoptionsTo.(*MapOptions)
		return
	}

	mapoptionsTo = new(MapOptions)
	mapOrigCopy[mapoptionsFrom] = mapoptionsTo
	mapoptionsFrom.CopyBasicFields(mapoptionsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layergroupuse := range mapoptionsFrom.LayerGroupUses {
		mapoptionsTo.LayerGroupUses = append( mapoptionsTo.LayerGroupUses, CopyBranchLayerGroupUse(mapOrigCopy, _layergroupuse))
	}

	return
}

func CopyBranchMarker(mapOrigCopy map[any]any, markerFrom *Marker) (markerTo  *Marker){

	// markerFrom has already been copied
	if _markerTo, ok := mapOrigCopy[markerFrom]; ok {
		markerTo = _markerTo.(*Marker)
		return
	}

	markerTo = new(Marker)
	mapOrigCopy[markerFrom] = markerTo
	markerFrom.CopyBasicFields(markerTo)

	//insertion point for the staging of instances referenced by pointers
	if markerFrom.LayerGroup != nil {
		markerTo.LayerGroup = CopyBranchLayerGroup(mapOrigCopy, markerFrom.LayerGroup)
	}
	if markerFrom.DivIcon != nil {
		markerTo.DivIcon = CopyBranchDivIcon(mapOrigCopy, markerFrom.DivIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUserClick(mapOrigCopy map[any]any, userclickFrom *UserClick) (userclickTo  *UserClick){

	// userclickFrom has already been copied
	if _userclickTo, ok := mapOrigCopy[userclickFrom]; ok {
		userclickTo = _userclickTo.(*UserClick)
		return
	}

	userclickTo = new(UserClick)
	mapOrigCopy[userclickFrom] = userclickTo
	userclickFrom.CopyBasicFields(userclickTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVLine(mapOrigCopy map[any]any, vlineFrom *VLine) (vlineTo  *VLine){

	// vlineFrom has already been copied
	if _vlineTo, ok := mapOrigCopy[vlineFrom]; ok {
		vlineTo = _vlineTo.(*VLine)
		return
	}

	vlineTo = new(VLine)
	mapOrigCopy[vlineFrom] = vlineTo
	vlineFrom.CopyBasicFields(vlineTo)

	//insertion point for the staging of instances referenced by pointers
	if vlineFrom.LayerGroup != nil {
		vlineTo.LayerGroup = CopyBranchLayerGroup(mapOrigCopy, vlineFrom.LayerGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVisualTrack(mapOrigCopy map[any]any, visualtrackFrom *VisualTrack) (visualtrackTo  *VisualTrack){

	// visualtrackFrom has already been copied
	if _visualtrackTo, ok := mapOrigCopy[visualtrackFrom]; ok {
		visualtrackTo = _visualtrackTo.(*VisualTrack)
		return
	}

	visualtrackTo = new(VisualTrack)
	mapOrigCopy[visualtrackFrom] = visualtrackTo
	visualtrackFrom.CopyBasicFields(visualtrackTo)

	//insertion point for the staging of instances referenced by pointers
	if visualtrackFrom.LayerGroup != nil {
		visualtrackTo.LayerGroup = CopyBranchLayerGroup(mapOrigCopy, visualtrackFrom.LayerGroup)
	}
	if visualtrackFrom.DivIcon != nil {
		visualtrackTo.DivIcon = CopyBranchDivIcon(mapOrigCopy, visualtrackFrom.DivIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
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

