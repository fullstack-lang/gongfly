package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CheckoutScheduler:
		if stage.OnAfterCheckoutSchedulerCreateCallback != nil {
			stage.OnAfterCheckoutSchedulerCreateCallback.OnAfterCreate(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleCreateCallback != nil {
			stage.OnAfterCircleCreateCallback.OnAfterCreate(stage, target)
		}
	case *DivIcon:
		if stage.OnAfterDivIconCreateCallback != nil {
			stage.OnAfterDivIconCreateCallback.OnAfterCreate(stage, target)
		}
	case *LayerGroup:
		if stage.OnAfterLayerGroupCreateCallback != nil {
			stage.OnAfterLayerGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *LayerGroupUse:
		if stage.OnAfterLayerGroupUseCreateCallback != nil {
			stage.OnAfterLayerGroupUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *MapOptions:
		if stage.OnAfterMapOptionsCreateCallback != nil {
			stage.OnAfterMapOptionsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Marker:
		if stage.OnAfterMarkerCreateCallback != nil {
			stage.OnAfterMarkerCreateCallback.OnAfterCreate(stage, target)
		}
	case *UserClick:
		if stage.OnAfterUserClickCreateCallback != nil {
			stage.OnAfterUserClickCreateCallback.OnAfterCreate(stage, target)
		}
	case *VLine:
		if stage.OnAfterVLineCreateCallback != nil {
			stage.OnAfterVLineCreateCallback.OnAfterCreate(stage, target)
		}
	case *VisualTrack:
		if stage.OnAfterVisualTrackCreateCallback != nil {
			stage.OnAfterVisualTrackCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *CheckoutScheduler:
		newTarget := any(new).(*CheckoutScheduler)
		if stage.OnAfterCheckoutSchedulerUpdateCallback != nil {
			stage.OnAfterCheckoutSchedulerUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Circle:
		newTarget := any(new).(*Circle)
		if stage.OnAfterCircleUpdateCallback != nil {
			stage.OnAfterCircleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DivIcon:
		newTarget := any(new).(*DivIcon)
		if stage.OnAfterDivIconUpdateCallback != nil {
			stage.OnAfterDivIconUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LayerGroup:
		newTarget := any(new).(*LayerGroup)
		if stage.OnAfterLayerGroupUpdateCallback != nil {
			stage.OnAfterLayerGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LayerGroupUse:
		newTarget := any(new).(*LayerGroupUse)
		if stage.OnAfterLayerGroupUseUpdateCallback != nil {
			stage.OnAfterLayerGroupUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MapOptions:
		newTarget := any(new).(*MapOptions)
		if stage.OnAfterMapOptionsUpdateCallback != nil {
			stage.OnAfterMapOptionsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Marker:
		newTarget := any(new).(*Marker)
		if stage.OnAfterMarkerUpdateCallback != nil {
			stage.OnAfterMarkerUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *UserClick:
		newTarget := any(new).(*UserClick)
		if stage.OnAfterUserClickUpdateCallback != nil {
			stage.OnAfterUserClickUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VLine:
		newTarget := any(new).(*VLine)
		if stage.OnAfterVLineUpdateCallback != nil {
			stage.OnAfterVLineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VisualTrack:
		newTarget := any(new).(*VisualTrack)
		if stage.OnAfterVisualTrackUpdateCallback != nil {
			stage.OnAfterVisualTrackUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *CheckoutScheduler:
		if stage.OnAfterCheckoutSchedulerDeleteCallback != nil {
			staged := any(staged).(*CheckoutScheduler)
			stage.OnAfterCheckoutSchedulerDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Circle:
		if stage.OnAfterCircleDeleteCallback != nil {
			staged := any(staged).(*Circle)
			stage.OnAfterCircleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DivIcon:
		if stage.OnAfterDivIconDeleteCallback != nil {
			staged := any(staged).(*DivIcon)
			stage.OnAfterDivIconDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LayerGroup:
		if stage.OnAfterLayerGroupDeleteCallback != nil {
			staged := any(staged).(*LayerGroup)
			stage.OnAfterLayerGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LayerGroupUse:
		if stage.OnAfterLayerGroupUseDeleteCallback != nil {
			staged := any(staged).(*LayerGroupUse)
			stage.OnAfterLayerGroupUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MapOptions:
		if stage.OnAfterMapOptionsDeleteCallback != nil {
			staged := any(staged).(*MapOptions)
			stage.OnAfterMapOptionsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Marker:
		if stage.OnAfterMarkerDeleteCallback != nil {
			staged := any(staged).(*Marker)
			stage.OnAfterMarkerDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *UserClick:
		if stage.OnAfterUserClickDeleteCallback != nil {
			staged := any(staged).(*UserClick)
			stage.OnAfterUserClickDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VLine:
		if stage.OnAfterVLineDeleteCallback != nil {
			staged := any(staged).(*VLine)
			stage.OnAfterVLineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VisualTrack:
		if stage.OnAfterVisualTrackDeleteCallback != nil {
			staged := any(staged).(*VisualTrack)
			stage.OnAfterVisualTrackDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CheckoutScheduler:
		if stage.OnAfterCheckoutSchedulerReadCallback != nil {
			stage.OnAfterCheckoutSchedulerReadCallback.OnAfterRead(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleReadCallback != nil {
			stage.OnAfterCircleReadCallback.OnAfterRead(stage, target)
		}
	case *DivIcon:
		if stage.OnAfterDivIconReadCallback != nil {
			stage.OnAfterDivIconReadCallback.OnAfterRead(stage, target)
		}
	case *LayerGroup:
		if stage.OnAfterLayerGroupReadCallback != nil {
			stage.OnAfterLayerGroupReadCallback.OnAfterRead(stage, target)
		}
	case *LayerGroupUse:
		if stage.OnAfterLayerGroupUseReadCallback != nil {
			stage.OnAfterLayerGroupUseReadCallback.OnAfterRead(stage, target)
		}
	case *MapOptions:
		if stage.OnAfterMapOptionsReadCallback != nil {
			stage.OnAfterMapOptionsReadCallback.OnAfterRead(stage, target)
		}
	case *Marker:
		if stage.OnAfterMarkerReadCallback != nil {
			stage.OnAfterMarkerReadCallback.OnAfterRead(stage, target)
		}
	case *UserClick:
		if stage.OnAfterUserClickReadCallback != nil {
			stage.OnAfterUserClickReadCallback.OnAfterRead(stage, target)
		}
	case *VLine:
		if stage.OnAfterVLineReadCallback != nil {
			stage.OnAfterVLineReadCallback.OnAfterRead(stage, target)
		}
	case *VisualTrack:
		if stage.OnAfterVisualTrackReadCallback != nil {
			stage.OnAfterVisualTrackReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CheckoutScheduler:
		stage.OnAfterCheckoutSchedulerUpdateCallback = any(callback).(OnAfterUpdateInterface[CheckoutScheduler])
	
	case *Circle:
		stage.OnAfterCircleUpdateCallback = any(callback).(OnAfterUpdateInterface[Circle])
	
	case *DivIcon:
		stage.OnAfterDivIconUpdateCallback = any(callback).(OnAfterUpdateInterface[DivIcon])
	
	case *LayerGroup:
		stage.OnAfterLayerGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[LayerGroup])
	
	case *LayerGroupUse:
		stage.OnAfterLayerGroupUseUpdateCallback = any(callback).(OnAfterUpdateInterface[LayerGroupUse])
	
	case *MapOptions:
		stage.OnAfterMapOptionsUpdateCallback = any(callback).(OnAfterUpdateInterface[MapOptions])
	
	case *Marker:
		stage.OnAfterMarkerUpdateCallback = any(callback).(OnAfterUpdateInterface[Marker])
	
	case *UserClick:
		stage.OnAfterUserClickUpdateCallback = any(callback).(OnAfterUpdateInterface[UserClick])
	
	case *VLine:
		stage.OnAfterVLineUpdateCallback = any(callback).(OnAfterUpdateInterface[VLine])
	
	case *VisualTrack:
		stage.OnAfterVisualTrackUpdateCallback = any(callback).(OnAfterUpdateInterface[VisualTrack])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CheckoutScheduler:
		stage.OnAfterCheckoutSchedulerCreateCallback = any(callback).(OnAfterCreateInterface[CheckoutScheduler])
	
	case *Circle:
		stage.OnAfterCircleCreateCallback = any(callback).(OnAfterCreateInterface[Circle])
	
	case *DivIcon:
		stage.OnAfterDivIconCreateCallback = any(callback).(OnAfterCreateInterface[DivIcon])
	
	case *LayerGroup:
		stage.OnAfterLayerGroupCreateCallback = any(callback).(OnAfterCreateInterface[LayerGroup])
	
	case *LayerGroupUse:
		stage.OnAfterLayerGroupUseCreateCallback = any(callback).(OnAfterCreateInterface[LayerGroupUse])
	
	case *MapOptions:
		stage.OnAfterMapOptionsCreateCallback = any(callback).(OnAfterCreateInterface[MapOptions])
	
	case *Marker:
		stage.OnAfterMarkerCreateCallback = any(callback).(OnAfterCreateInterface[Marker])
	
	case *UserClick:
		stage.OnAfterUserClickCreateCallback = any(callback).(OnAfterCreateInterface[UserClick])
	
	case *VLine:
		stage.OnAfterVLineCreateCallback = any(callback).(OnAfterCreateInterface[VLine])
	
	case *VisualTrack:
		stage.OnAfterVisualTrackCreateCallback = any(callback).(OnAfterCreateInterface[VisualTrack])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CheckoutScheduler:
		stage.OnAfterCheckoutSchedulerDeleteCallback = any(callback).(OnAfterDeleteInterface[CheckoutScheduler])
	
	case *Circle:
		stage.OnAfterCircleDeleteCallback = any(callback).(OnAfterDeleteInterface[Circle])
	
	case *DivIcon:
		stage.OnAfterDivIconDeleteCallback = any(callback).(OnAfterDeleteInterface[DivIcon])
	
	case *LayerGroup:
		stage.OnAfterLayerGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[LayerGroup])
	
	case *LayerGroupUse:
		stage.OnAfterLayerGroupUseDeleteCallback = any(callback).(OnAfterDeleteInterface[LayerGroupUse])
	
	case *MapOptions:
		stage.OnAfterMapOptionsDeleteCallback = any(callback).(OnAfterDeleteInterface[MapOptions])
	
	case *Marker:
		stage.OnAfterMarkerDeleteCallback = any(callback).(OnAfterDeleteInterface[Marker])
	
	case *UserClick:
		stage.OnAfterUserClickDeleteCallback = any(callback).(OnAfterDeleteInterface[UserClick])
	
	case *VLine:
		stage.OnAfterVLineDeleteCallback = any(callback).(OnAfterDeleteInterface[VLine])
	
	case *VisualTrack:
		stage.OnAfterVisualTrackDeleteCallback = any(callback).(OnAfterDeleteInterface[VisualTrack])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CheckoutScheduler:
		stage.OnAfterCheckoutSchedulerReadCallback = any(callback).(OnAfterReadInterface[CheckoutScheduler])
	
	case *Circle:
		stage.OnAfterCircleReadCallback = any(callback).(OnAfterReadInterface[Circle])
	
	case *DivIcon:
		stage.OnAfterDivIconReadCallback = any(callback).(OnAfterReadInterface[DivIcon])
	
	case *LayerGroup:
		stage.OnAfterLayerGroupReadCallback = any(callback).(OnAfterReadInterface[LayerGroup])
	
	case *LayerGroupUse:
		stage.OnAfterLayerGroupUseReadCallback = any(callback).(OnAfterReadInterface[LayerGroupUse])
	
	case *MapOptions:
		stage.OnAfterMapOptionsReadCallback = any(callback).(OnAfterReadInterface[MapOptions])
	
	case *Marker:
		stage.OnAfterMarkerReadCallback = any(callback).(OnAfterReadInterface[Marker])
	
	case *UserClick:
		stage.OnAfterUserClickReadCallback = any(callback).(OnAfterReadInterface[UserClick])
	
	case *VLine:
		stage.OnAfterVLineReadCallback = any(callback).(OnAfterReadInterface[VLine])
	
	case *VisualTrack:
		stage.OnAfterVisualTrackReadCallback = any(callback).(OnAfterReadInterface[VisualTrack])
	
	}
}
