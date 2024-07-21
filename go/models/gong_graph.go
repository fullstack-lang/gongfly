// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CivilianAirport:
		ok = stage.IsStagedCivilianAirport(target)

	case *Liner:
		ok = stage.IsStagedLiner(target)

	case *Message:
		ok = stage.IsStagedMessage(target)

	case *OpsLine:
		ok = stage.IsStagedOpsLine(target)

	case *Radar:
		ok = stage.IsStagedRadar(target)

	case *Satellite:
		ok = stage.IsStagedSatellite(target)

	case *Scenario:
		ok = stage.IsStagedScenario(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedCivilianAirport(civilianairport *CivilianAirport) (ok bool) {

	_, ok = stage.CivilianAirports[civilianairport]

	return
}

func (stage *StageStruct) IsStagedLiner(liner *Liner) (ok bool) {

	_, ok = stage.Liners[liner]

	return
}

func (stage *StageStruct) IsStagedMessage(message *Message) (ok bool) {

	_, ok = stage.Messages[message]

	return
}

func (stage *StageStruct) IsStagedOpsLine(opsline *OpsLine) (ok bool) {

	_, ok = stage.OpsLines[opsline]

	return
}

func (stage *StageStruct) IsStagedRadar(radar *Radar) (ok bool) {

	_, ok = stage.Radars[radar]

	return
}

func (stage *StageStruct) IsStagedSatellite(satellite *Satellite) (ok bool) {

	_, ok = stage.Satellites[satellite]

	return
}

func (stage *StageStruct) IsStagedScenario(scenario *Scenario) (ok bool) {

	_, ok = stage.Scenarios[scenario]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *CivilianAirport:
		stage.StageBranchCivilianAirport(target)

	case *Liner:
		stage.StageBranchLiner(target)

	case *Message:
		stage.StageBranchMessage(target)

	case *OpsLine:
		stage.StageBranchOpsLine(target)

	case *Radar:
		stage.StageBranchRadar(target)

	case *Satellite:
		stage.StageBranchSatellite(target)

	case *Scenario:
		stage.StageBranchScenario(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchCivilianAirport(civilianairport *CivilianAirport) {

	// check if instance is already staged
	if IsStaged(stage, civilianairport) {
		return
	}

	civilianairport.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLiner(liner *Liner) {

	// check if instance is already staged
	if IsStaged(stage, liner) {
		return
	}

	liner.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if liner.ReporingLine != nil {
		StageBranch(stage, liner.ReporingLine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMessage(message *Message) {

	// check if instance is already staged
	if IsStaged(stage, message) {
		return
	}

	message.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOpsLine(opsline *OpsLine) {

	// check if instance is already staged
	if IsStaged(stage, opsline) {
		return
	}

	opsline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if opsline.Scenario != nil {
		StageBranch(stage, opsline.Scenario)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRadar(radar *Radar) {

	// check if instance is already staged
	if IsStaged(stage, radar) {
		return
	}

	radar.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSatellite(satellite *Satellite) {

	// check if instance is already staged
	if IsStaged(stage, satellite) {
		return
	}

	satellite.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchScenario(scenario *Scenario) {

	// check if instance is already staged
	if IsStaged(stage, scenario) {
		return
	}

	scenario.Stage(stage)

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
	case *CivilianAirport:
		toT := CopyBranchCivilianAirport(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Liner:
		toT := CopyBranchLiner(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Message:
		toT := CopyBranchMessage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OpsLine:
		toT := CopyBranchOpsLine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Radar:
		toT := CopyBranchRadar(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Satellite:
		toT := CopyBranchSatellite(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scenario:
		toT := CopyBranchScenario(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCivilianAirport(mapOrigCopy map[any]any, civilianairportFrom *CivilianAirport) (civilianairportTo *CivilianAirport) {

	// civilianairportFrom has already been copied
	if _civilianairportTo, ok := mapOrigCopy[civilianairportFrom]; ok {
		civilianairportTo = _civilianairportTo.(*CivilianAirport)
		return
	}

	civilianairportTo = new(CivilianAirport)
	mapOrigCopy[civilianairportFrom] = civilianairportTo
	civilianairportFrom.CopyBasicFields(civilianairportTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLiner(mapOrigCopy map[any]any, linerFrom *Liner) (linerTo *Liner) {

	// linerFrom has already been copied
	if _linerTo, ok := mapOrigCopy[linerFrom]; ok {
		linerTo = _linerTo.(*Liner)
		return
	}

	linerTo = new(Liner)
	mapOrigCopy[linerFrom] = linerTo
	linerFrom.CopyBasicFields(linerTo)

	//insertion point for the staging of instances referenced by pointers
	if linerFrom.ReporingLine != nil {
		linerTo.ReporingLine = CopyBranchOpsLine(mapOrigCopy, linerFrom.ReporingLine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMessage(mapOrigCopy map[any]any, messageFrom *Message) (messageTo *Message) {

	// messageFrom has already been copied
	if _messageTo, ok := mapOrigCopy[messageFrom]; ok {
		messageTo = _messageTo.(*Message)
		return
	}

	messageTo = new(Message)
	mapOrigCopy[messageFrom] = messageTo
	messageFrom.CopyBasicFields(messageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOpsLine(mapOrigCopy map[any]any, opslineFrom *OpsLine) (opslineTo *OpsLine) {

	// opslineFrom has already been copied
	if _opslineTo, ok := mapOrigCopy[opslineFrom]; ok {
		opslineTo = _opslineTo.(*OpsLine)
		return
	}

	opslineTo = new(OpsLine)
	mapOrigCopy[opslineFrom] = opslineTo
	opslineFrom.CopyBasicFields(opslineTo)

	//insertion point for the staging of instances referenced by pointers
	if opslineFrom.Scenario != nil {
		opslineTo.Scenario = CopyBranchScenario(mapOrigCopy, opslineFrom.Scenario)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRadar(mapOrigCopy map[any]any, radarFrom *Radar) (radarTo *Radar) {

	// radarFrom has already been copied
	if _radarTo, ok := mapOrigCopy[radarFrom]; ok {
		radarTo = _radarTo.(*Radar)
		return
	}

	radarTo = new(Radar)
	mapOrigCopy[radarFrom] = radarTo
	radarFrom.CopyBasicFields(radarTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSatellite(mapOrigCopy map[any]any, satelliteFrom *Satellite) (satelliteTo *Satellite) {

	// satelliteFrom has already been copied
	if _satelliteTo, ok := mapOrigCopy[satelliteFrom]; ok {
		satelliteTo = _satelliteTo.(*Satellite)
		return
	}

	satelliteTo = new(Satellite)
	mapOrigCopy[satelliteFrom] = satelliteTo
	satelliteFrom.CopyBasicFields(satelliteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchScenario(mapOrigCopy map[any]any, scenarioFrom *Scenario) (scenarioTo *Scenario) {

	// scenarioFrom has already been copied
	if _scenarioTo, ok := mapOrigCopy[scenarioFrom]; ok {
		scenarioTo = _scenarioTo.(*Scenario)
		return
	}

	scenarioTo = new(Scenario)
	mapOrigCopy[scenarioFrom] = scenarioTo
	scenarioFrom.CopyBasicFields(scenarioTo)

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
	case *CivilianAirport:
		stage.UnstageBranchCivilianAirport(target)

	case *Liner:
		stage.UnstageBranchLiner(target)

	case *Message:
		stage.UnstageBranchMessage(target)

	case *OpsLine:
		stage.UnstageBranchOpsLine(target)

	case *Radar:
		stage.UnstageBranchRadar(target)

	case *Satellite:
		stage.UnstageBranchSatellite(target)

	case *Scenario:
		stage.UnstageBranchScenario(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchCivilianAirport(civilianairport *CivilianAirport) {

	// check if instance is already staged
	if !IsStaged(stage, civilianairport) {
		return
	}

	civilianairport.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLiner(liner *Liner) {

	// check if instance is already staged
	if !IsStaged(stage, liner) {
		return
	}

	liner.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if liner.ReporingLine != nil {
		UnstageBranch(stage, liner.ReporingLine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMessage(message *Message) {

	// check if instance is already staged
	if !IsStaged(stage, message) {
		return
	}

	message.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOpsLine(opsline *OpsLine) {

	// check if instance is already staged
	if !IsStaged(stage, opsline) {
		return
	}

	opsline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if opsline.Scenario != nil {
		UnstageBranch(stage, opsline.Scenario)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRadar(radar *Radar) {

	// check if instance is already staged
	if !IsStaged(stage, radar) {
		return
	}

	radar.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSatellite(satellite *Satellite) {

	// check if instance is already staged
	if !IsStaged(stage, satellite) {
		return
	}

	satellite.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchScenario(scenario *Scenario) {

	// check if instance is already staged
	if !IsStaged(stage, scenario) {
		return
	}

	scenario.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
