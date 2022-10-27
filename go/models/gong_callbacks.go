package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CivilianAirport:
		if stage.OnAfterCivilianAirportCreateCallback != nil {
			stage.OnAfterCivilianAirportCreateCallback.OnAfterCreate(stage, target)
		}
	case *Liner:
		if stage.OnAfterLinerCreateCallback != nil {
			stage.OnAfterLinerCreateCallback.OnAfterCreate(stage, target)
		}
	case *Message:
		if stage.OnAfterMessageCreateCallback != nil {
			stage.OnAfterMessageCreateCallback.OnAfterCreate(stage, target)
		}
	case *OpsLine:
		if stage.OnAfterOpsLineCreateCallback != nil {
			stage.OnAfterOpsLineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Radar:
		if stage.OnAfterRadarCreateCallback != nil {
			stage.OnAfterRadarCreateCallback.OnAfterCreate(stage, target)
		}
	case *Satellite:
		if stage.OnAfterSatelliteCreateCallback != nil {
			stage.OnAfterSatelliteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Scenario:
		if stage.OnAfterScenarioCreateCallback != nil {
			stage.OnAfterScenarioCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *CivilianAirport:
		newTarget := any(new).(*CivilianAirport)
		if stage.OnAfterCivilianAirportUpdateCallback != nil {
			stage.OnAfterCivilianAirportUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Liner:
		newTarget := any(new).(*Liner)
		if stage.OnAfterLinerUpdateCallback != nil {
			stage.OnAfterLinerUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Message:
		newTarget := any(new).(*Message)
		if stage.OnAfterMessageUpdateCallback != nil {
			stage.OnAfterMessageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *OpsLine:
		newTarget := any(new).(*OpsLine)
		if stage.OnAfterOpsLineUpdateCallback != nil {
			stage.OnAfterOpsLineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Radar:
		newTarget := any(new).(*Radar)
		if stage.OnAfterRadarUpdateCallback != nil {
			stage.OnAfterRadarUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Satellite:
		newTarget := any(new).(*Satellite)
		if stage.OnAfterSatelliteUpdateCallback != nil {
			stage.OnAfterSatelliteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Scenario:
		newTarget := any(new).(*Scenario)
		if stage.OnAfterScenarioUpdateCallback != nil {
			stage.OnAfterScenarioUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *CivilianAirport:
		if stage.OnAfterCivilianAirportDeleteCallback != nil {
			staged := any(staged).(*CivilianAirport)
			stage.OnAfterCivilianAirportDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Liner:
		if stage.OnAfterLinerDeleteCallback != nil {
			staged := any(staged).(*Liner)
			stage.OnAfterLinerDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Message:
		if stage.OnAfterMessageDeleteCallback != nil {
			staged := any(staged).(*Message)
			stage.OnAfterMessageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *OpsLine:
		if stage.OnAfterOpsLineDeleteCallback != nil {
			staged := any(staged).(*OpsLine)
			stage.OnAfterOpsLineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Radar:
		if stage.OnAfterRadarDeleteCallback != nil {
			staged := any(staged).(*Radar)
			stage.OnAfterRadarDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Satellite:
		if stage.OnAfterSatelliteDeleteCallback != nil {
			staged := any(staged).(*Satellite)
			stage.OnAfterSatelliteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Scenario:
		if stage.OnAfterScenarioDeleteCallback != nil {
			staged := any(staged).(*Scenario)
			stage.OnAfterScenarioDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CivilianAirport:
		if stage.OnAfterCivilianAirportReadCallback != nil {
			stage.OnAfterCivilianAirportReadCallback.OnAfterRead(stage, target)
		}
	case *Liner:
		if stage.OnAfterLinerReadCallback != nil {
			stage.OnAfterLinerReadCallback.OnAfterRead(stage, target)
		}
	case *Message:
		if stage.OnAfterMessageReadCallback != nil {
			stage.OnAfterMessageReadCallback.OnAfterRead(stage, target)
		}
	case *OpsLine:
		if stage.OnAfterOpsLineReadCallback != nil {
			stage.OnAfterOpsLineReadCallback.OnAfterRead(stage, target)
		}
	case *Radar:
		if stage.OnAfterRadarReadCallback != nil {
			stage.OnAfterRadarReadCallback.OnAfterRead(stage, target)
		}
	case *Satellite:
		if stage.OnAfterSatelliteReadCallback != nil {
			stage.OnAfterSatelliteReadCallback.OnAfterRead(stage, target)
		}
	case *Scenario:
		if stage.OnAfterScenarioReadCallback != nil {
			stage.OnAfterScenarioReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CivilianAirport:
		stage.OnAfterCivilianAirportUpdateCallback = any(callback).(OnAfterUpdateInterface[CivilianAirport])
	
	case *Liner:
		stage.OnAfterLinerUpdateCallback = any(callback).(OnAfterUpdateInterface[Liner])
	
	case *Message:
		stage.OnAfterMessageUpdateCallback = any(callback).(OnAfterUpdateInterface[Message])
	
	case *OpsLine:
		stage.OnAfterOpsLineUpdateCallback = any(callback).(OnAfterUpdateInterface[OpsLine])
	
	case *Radar:
		stage.OnAfterRadarUpdateCallback = any(callback).(OnAfterUpdateInterface[Radar])
	
	case *Satellite:
		stage.OnAfterSatelliteUpdateCallback = any(callback).(OnAfterUpdateInterface[Satellite])
	
	case *Scenario:
		stage.OnAfterScenarioUpdateCallback = any(callback).(OnAfterUpdateInterface[Scenario])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CivilianAirport:
		stage.OnAfterCivilianAirportCreateCallback = any(callback).(OnAfterCreateInterface[CivilianAirport])
	
	case *Liner:
		stage.OnAfterLinerCreateCallback = any(callback).(OnAfterCreateInterface[Liner])
	
	case *Message:
		stage.OnAfterMessageCreateCallback = any(callback).(OnAfterCreateInterface[Message])
	
	case *OpsLine:
		stage.OnAfterOpsLineCreateCallback = any(callback).(OnAfterCreateInterface[OpsLine])
	
	case *Radar:
		stage.OnAfterRadarCreateCallback = any(callback).(OnAfterCreateInterface[Radar])
	
	case *Satellite:
		stage.OnAfterSatelliteCreateCallback = any(callback).(OnAfterCreateInterface[Satellite])
	
	case *Scenario:
		stage.OnAfterScenarioCreateCallback = any(callback).(OnAfterCreateInterface[Scenario])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CivilianAirport:
		stage.OnAfterCivilianAirportDeleteCallback = any(callback).(OnAfterDeleteInterface[CivilianAirport])
	
	case *Liner:
		stage.OnAfterLinerDeleteCallback = any(callback).(OnAfterDeleteInterface[Liner])
	
	case *Message:
		stage.OnAfterMessageDeleteCallback = any(callback).(OnAfterDeleteInterface[Message])
	
	case *OpsLine:
		stage.OnAfterOpsLineDeleteCallback = any(callback).(OnAfterDeleteInterface[OpsLine])
	
	case *Radar:
		stage.OnAfterRadarDeleteCallback = any(callback).(OnAfterDeleteInterface[Radar])
	
	case *Satellite:
		stage.OnAfterSatelliteDeleteCallback = any(callback).(OnAfterDeleteInterface[Satellite])
	
	case *Scenario:
		stage.OnAfterScenarioDeleteCallback = any(callback).(OnAfterDeleteInterface[Scenario])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CivilianAirport:
		stage.OnAfterCivilianAirportReadCallback = any(callback).(OnAfterReadInterface[CivilianAirport])
	
	case *Liner:
		stage.OnAfterLinerReadCallback = any(callback).(OnAfterReadInterface[Liner])
	
	case *Message:
		stage.OnAfterMessageReadCallback = any(callback).(OnAfterReadInterface[Message])
	
	case *OpsLine:
		stage.OnAfterOpsLineReadCallback = any(callback).(OnAfterReadInterface[OpsLine])
	
	case *Radar:
		stage.OnAfterRadarReadCallback = any(callback).(OnAfterReadInterface[Radar])
	
	case *Satellite:
		stage.OnAfterSatelliteReadCallback = any(callback).(OnAfterReadInterface[Satellite])
	
	case *Scenario:
		stage.OnAfterScenarioReadCallback = any(callback).(OnAfterReadInterface[Scenario])
	
	}
}
