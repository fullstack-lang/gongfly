// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *DummyAgent:
		ok = stage.IsStagedDummyAgent(target)

	case *Engine:
		ok = stage.IsStagedEngine(target)

	case *Event:
		ok = stage.IsStagedEvent(target)

	case *GongsimCommand:
		ok = stage.IsStagedGongsimCommand(target)

	case *GongsimStatus:
		ok = stage.IsStagedGongsimStatus(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedDummyAgent(dummyagent *DummyAgent) (ok bool) {

	_, ok = stage.DummyAgents[dummyagent]

	return
}

func (stage *StageStruct) IsStagedEngine(engine *Engine) (ok bool) {

	_, ok = stage.Engines[engine]

	return
}

func (stage *StageStruct) IsStagedEvent(event *Event) (ok bool) {

	_, ok = stage.Events[event]

	return
}

func (stage *StageStruct) IsStagedGongsimCommand(gongsimcommand *GongsimCommand) (ok bool) {

	_, ok = stage.GongsimCommands[gongsimcommand]

	return
}

func (stage *StageStruct) IsStagedGongsimStatus(gongsimstatus *GongsimStatus) (ok bool) {

	_, ok = stage.GongsimStatuss[gongsimstatus]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *DummyAgent:
		stage.StageBranchDummyAgent(target)

	case *Engine:
		stage.StageBranchEngine(target)

	case *Event:
		stage.StageBranchEvent(target)

	case *GongsimCommand:
		stage.StageBranchGongsimCommand(target)

	case *GongsimStatus:
		stage.StageBranchGongsimStatus(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchDummyAgent(dummyagent *DummyAgent) {

	// check if instance is already staged
	if IsStaged(stage, dummyagent) {
		return
	}

	dummyagent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEngine(engine *Engine) {

	// check if instance is already staged
	if IsStaged(stage, engine) {
		return
	}

	engine.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEvent(event *Event) {

	// check if instance is already staged
	if IsStaged(stage, event) {
		return
	}

	event.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongsimCommand(gongsimcommand *GongsimCommand) {

	// check if instance is already staged
	if IsStaged(stage, gongsimcommand) {
		return
	}

	gongsimcommand.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongsimcommand.Engine != nil {
		StageBranch(stage, gongsimcommand.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongsimStatus(gongsimstatus *GongsimStatus) {

	// check if instance is already staged
	if IsStaged(stage, gongsimstatus) {
		return
	}

	gongsimstatus.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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
	case *DummyAgent:
		toT := CopyBranchDummyAgent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Engine:
		toT := CopyBranchEngine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Event:
		toT := CopyBranchEvent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongsimCommand:
		toT := CopyBranchGongsimCommand(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongsimStatus:
		toT := CopyBranchGongsimStatus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchDummyAgent(mapOrigCopy map[any]any, dummyagentFrom *DummyAgent) (dummyagentTo *DummyAgent) {

	// dummyagentFrom has already been copied
	if _dummyagentTo, ok := mapOrigCopy[dummyagentFrom]; ok {
		dummyagentTo = _dummyagentTo.(*DummyAgent)
		return
	}

	dummyagentTo = new(DummyAgent)
	mapOrigCopy[dummyagentFrom] = dummyagentTo
	dummyagentFrom.CopyBasicFields(dummyagentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEngine(mapOrigCopy map[any]any, engineFrom *Engine) (engineTo *Engine) {

	// engineFrom has already been copied
	if _engineTo, ok := mapOrigCopy[engineFrom]; ok {
		engineTo = _engineTo.(*Engine)
		return
	}

	engineTo = new(Engine)
	mapOrigCopy[engineFrom] = engineTo
	engineFrom.CopyBasicFields(engineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEvent(mapOrigCopy map[any]any, eventFrom *Event) (eventTo *Event) {

	// eventFrom has already been copied
	if _eventTo, ok := mapOrigCopy[eventFrom]; ok {
		eventTo = _eventTo.(*Event)
		return
	}

	eventTo = new(Event)
	mapOrigCopy[eventFrom] = eventTo
	eventFrom.CopyBasicFields(eventTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGongsimCommand(mapOrigCopy map[any]any, gongsimcommandFrom *GongsimCommand) (gongsimcommandTo *GongsimCommand) {

	// gongsimcommandFrom has already been copied
	if _gongsimcommandTo, ok := mapOrigCopy[gongsimcommandFrom]; ok {
		gongsimcommandTo = _gongsimcommandTo.(*GongsimCommand)
		return
	}

	gongsimcommandTo = new(GongsimCommand)
	mapOrigCopy[gongsimcommandFrom] = gongsimcommandTo
	gongsimcommandFrom.CopyBasicFields(gongsimcommandTo)

	//insertion point for the staging of instances referenced by pointers
	if gongsimcommandFrom.Engine != nil {
		gongsimcommandTo.Engine = CopyBranchEngine(mapOrigCopy, gongsimcommandFrom.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGongsimStatus(mapOrigCopy map[any]any, gongsimstatusFrom *GongsimStatus) (gongsimstatusTo *GongsimStatus) {

	// gongsimstatusFrom has already been copied
	if _gongsimstatusTo, ok := mapOrigCopy[gongsimstatusFrom]; ok {
		gongsimstatusTo = _gongsimstatusTo.(*GongsimStatus)
		return
	}

	gongsimstatusTo = new(GongsimStatus)
	mapOrigCopy[gongsimstatusFrom] = gongsimstatusTo
	gongsimstatusFrom.CopyBasicFields(gongsimstatusTo)

	//insertion point for the staging of instances referenced by pointers

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
	case *DummyAgent:
		stage.UnstageBranchDummyAgent(target)

	case *Engine:
		stage.UnstageBranchEngine(target)

	case *Event:
		stage.UnstageBranchEvent(target)

	case *GongsimCommand:
		stage.UnstageBranchGongsimCommand(target)

	case *GongsimStatus:
		stage.UnstageBranchGongsimStatus(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchDummyAgent(dummyagent *DummyAgent) {

	// check if instance is already staged
	if !IsStaged(stage, dummyagent) {
		return
	}

	dummyagent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEngine(engine *Engine) {

	// check if instance is already staged
	if !IsStaged(stage, engine) {
		return
	}

	engine.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEvent(event *Event) {

	// check if instance is already staged
	if !IsStaged(stage, event) {
		return
	}

	event.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongsimCommand(gongsimcommand *GongsimCommand) {

	// check if instance is already staged
	if !IsStaged(stage, gongsimcommand) {
		return
	}

	gongsimcommand.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongsimcommand.Engine != nil {
		UnstageBranch(stage, gongsimcommand.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongsimStatus(gongsimstatus *GongsimStatus) {

	// check if instance is already staged
	if !IsStaged(stage, gongsimstatus) {
		return
	}

	gongsimstatus.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
