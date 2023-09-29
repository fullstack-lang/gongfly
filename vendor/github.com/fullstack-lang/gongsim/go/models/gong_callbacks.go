// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DummyAgent:
		if stage.OnAfterDummyAgentCreateCallback != nil {
			stage.OnAfterDummyAgentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Engine:
		if stage.OnAfterEngineCreateCallback != nil {
			stage.OnAfterEngineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Event:
		if stage.OnAfterEventCreateCallback != nil {
			stage.OnAfterEventCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongsimCommand:
		if stage.OnAfterGongsimCommandCreateCallback != nil {
			stage.OnAfterGongsimCommandCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongsimStatus:
		if stage.OnAfterGongsimStatusCreateCallback != nil {
			stage.OnAfterGongsimStatusCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *DummyAgent:
		newTarget := any(new).(*DummyAgent)
		if stage.OnAfterDummyAgentUpdateCallback != nil {
			stage.OnAfterDummyAgentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Engine:
		newTarget := any(new).(*Engine)
		if stage.OnAfterEngineUpdateCallback != nil {
			stage.OnAfterEngineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Event:
		newTarget := any(new).(*Event)
		if stage.OnAfterEventUpdateCallback != nil {
			stage.OnAfterEventUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongsimCommand:
		newTarget := any(new).(*GongsimCommand)
		if stage.OnAfterGongsimCommandUpdateCallback != nil {
			stage.OnAfterGongsimCommandUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongsimStatus:
		newTarget := any(new).(*GongsimStatus)
		if stage.OnAfterGongsimStatusUpdateCallback != nil {
			stage.OnAfterGongsimStatusUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *DummyAgent:
		if stage.OnAfterDummyAgentDeleteCallback != nil {
			staged := any(staged).(*DummyAgent)
			stage.OnAfterDummyAgentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Engine:
		if stage.OnAfterEngineDeleteCallback != nil {
			staged := any(staged).(*Engine)
			stage.OnAfterEngineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Event:
		if stage.OnAfterEventDeleteCallback != nil {
			staged := any(staged).(*Event)
			stage.OnAfterEventDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongsimCommand:
		if stage.OnAfterGongsimCommandDeleteCallback != nil {
			staged := any(staged).(*GongsimCommand)
			stage.OnAfterGongsimCommandDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongsimStatus:
		if stage.OnAfterGongsimStatusDeleteCallback != nil {
			staged := any(staged).(*GongsimStatus)
			stage.OnAfterGongsimStatusDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DummyAgent:
		if stage.OnAfterDummyAgentReadCallback != nil {
			stage.OnAfterDummyAgentReadCallback.OnAfterRead(stage, target)
		}
	case *Engine:
		if stage.OnAfterEngineReadCallback != nil {
			stage.OnAfterEngineReadCallback.OnAfterRead(stage, target)
		}
	case *Event:
		if stage.OnAfterEventReadCallback != nil {
			stage.OnAfterEventReadCallback.OnAfterRead(stage, target)
		}
	case *GongsimCommand:
		if stage.OnAfterGongsimCommandReadCallback != nil {
			stage.OnAfterGongsimCommandReadCallback.OnAfterRead(stage, target)
		}
	case *GongsimStatus:
		if stage.OnAfterGongsimStatusReadCallback != nil {
			stage.OnAfterGongsimStatusReadCallback.OnAfterRead(stage, target)
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
	case *DummyAgent:
		stage.OnAfterDummyAgentUpdateCallback = any(callback).(OnAfterUpdateInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineUpdateCallback = any(callback).(OnAfterUpdateInterface[Engine])
	
	case *Event:
		stage.OnAfterEventUpdateCallback = any(callback).(OnAfterUpdateInterface[Event])
	
	case *GongsimCommand:
		stage.OnAfterGongsimCommandUpdateCallback = any(callback).(OnAfterUpdateInterface[GongsimCommand])
	
	case *GongsimStatus:
		stage.OnAfterGongsimStatusUpdateCallback = any(callback).(OnAfterUpdateInterface[GongsimStatus])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DummyAgent:
		stage.OnAfterDummyAgentCreateCallback = any(callback).(OnAfterCreateInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineCreateCallback = any(callback).(OnAfterCreateInterface[Engine])
	
	case *Event:
		stage.OnAfterEventCreateCallback = any(callback).(OnAfterCreateInterface[Event])
	
	case *GongsimCommand:
		stage.OnAfterGongsimCommandCreateCallback = any(callback).(OnAfterCreateInterface[GongsimCommand])
	
	case *GongsimStatus:
		stage.OnAfterGongsimStatusCreateCallback = any(callback).(OnAfterCreateInterface[GongsimStatus])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DummyAgent:
		stage.OnAfterDummyAgentDeleteCallback = any(callback).(OnAfterDeleteInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineDeleteCallback = any(callback).(OnAfterDeleteInterface[Engine])
	
	case *Event:
		stage.OnAfterEventDeleteCallback = any(callback).(OnAfterDeleteInterface[Event])
	
	case *GongsimCommand:
		stage.OnAfterGongsimCommandDeleteCallback = any(callback).(OnAfterDeleteInterface[GongsimCommand])
	
	case *GongsimStatus:
		stage.OnAfterGongsimStatusDeleteCallback = any(callback).(OnAfterDeleteInterface[GongsimStatus])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DummyAgent:
		stage.OnAfterDummyAgentReadCallback = any(callback).(OnAfterReadInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineReadCallback = any(callback).(OnAfterReadInterface[Engine])
	
	case *Event:
		stage.OnAfterEventReadCallback = any(callback).(OnAfterReadInterface[Event])
	
	case *GongsimCommand:
		stage.OnAfterGongsimCommandReadCallback = any(callback).(OnAfterReadInterface[GongsimCommand])
	
	case *GongsimStatus:
		stage.OnAfterGongsimStatusReadCallback = any(callback).(OnAfterReadInterface[GongsimStatus])
	
	}
}
