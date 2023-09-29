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
// referenced by pointers or slices of pointers of the insance
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
	if ! IsStaged(stage, civilianairport) {
		return
	}

	civilianairport.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLiner(liner *Liner) {

	// check if instance is already staged
	if ! IsStaged(stage, liner) {
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
	if ! IsStaged(stage, message) {
		return
	}

	message.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOpsLine(opsline *OpsLine) {

	// check if instance is already staged
	if ! IsStaged(stage, opsline) {
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
	if ! IsStaged(stage, radar) {
		return
	}

	radar.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSatellite(satellite *Satellite) {

	// check if instance is already staged
	if ! IsStaged(stage, satellite) {
		return
	}

	satellite.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchScenario(scenario *Scenario) {

	// check if instance is already staged
	if ! IsStaged(stage, scenario) {
		return
	}

	scenario.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

