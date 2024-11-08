// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	DummyAgentAPIs []*DummyAgentAPI

	EngineAPIs []*EngineAPI

	EventAPIs []*EventAPI

	GongsimCommandAPIs []*GongsimCommandAPI

	GongsimStatusAPIs []*GongsimStatusAPI

	UpdateStateAPIs []*UpdateStateAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, dummyagentDB := range backRepo.BackRepoDummyAgent.Map_DummyAgentDBID_DummyAgentDB {

		var dummyagentAPI DummyAgentAPI
		dummyagentAPI.ID = dummyagentDB.ID
		dummyagentAPI.DummyAgentPointersEncoding = dummyagentDB.DummyAgentPointersEncoding
		dummyagentDB.CopyBasicFieldsToDummyAgent_WOP(&dummyagentAPI.DummyAgent_WOP)

		backRepoData.DummyAgentAPIs = append(backRepoData.DummyAgentAPIs, &dummyagentAPI)
	}

	for _, engineDB := range backRepo.BackRepoEngine.Map_EngineDBID_EngineDB {

		var engineAPI EngineAPI
		engineAPI.ID = engineDB.ID
		engineAPI.EnginePointersEncoding = engineDB.EnginePointersEncoding
		engineDB.CopyBasicFieldsToEngine_WOP(&engineAPI.Engine_WOP)

		backRepoData.EngineAPIs = append(backRepoData.EngineAPIs, &engineAPI)
	}

	for _, eventDB := range backRepo.BackRepoEvent.Map_EventDBID_EventDB {

		var eventAPI EventAPI
		eventAPI.ID = eventDB.ID
		eventAPI.EventPointersEncoding = eventDB.EventPointersEncoding
		eventDB.CopyBasicFieldsToEvent_WOP(&eventAPI.Event_WOP)

		backRepoData.EventAPIs = append(backRepoData.EventAPIs, &eventAPI)
	}

	for _, gongsimcommandDB := range backRepo.BackRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB {

		var gongsimcommandAPI GongsimCommandAPI
		gongsimcommandAPI.ID = gongsimcommandDB.ID
		gongsimcommandAPI.GongsimCommandPointersEncoding = gongsimcommandDB.GongsimCommandPointersEncoding
		gongsimcommandDB.CopyBasicFieldsToGongsimCommand_WOP(&gongsimcommandAPI.GongsimCommand_WOP)

		backRepoData.GongsimCommandAPIs = append(backRepoData.GongsimCommandAPIs, &gongsimcommandAPI)
	}

	for _, gongsimstatusDB := range backRepo.BackRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB {

		var gongsimstatusAPI GongsimStatusAPI
		gongsimstatusAPI.ID = gongsimstatusDB.ID
		gongsimstatusAPI.GongsimStatusPointersEncoding = gongsimstatusDB.GongsimStatusPointersEncoding
		gongsimstatusDB.CopyBasicFieldsToGongsimStatus_WOP(&gongsimstatusAPI.GongsimStatus_WOP)

		backRepoData.GongsimStatusAPIs = append(backRepoData.GongsimStatusAPIs, &gongsimstatusAPI)
	}

	for _, updatestateDB := range backRepo.BackRepoUpdateState.Map_UpdateStateDBID_UpdateStateDB {

		var updatestateAPI UpdateStateAPI
		updatestateAPI.ID = updatestateDB.ID
		updatestateAPI.UpdateStatePointersEncoding = updatestateDB.UpdateStatePointersEncoding
		updatestateDB.CopyBasicFieldsToUpdateState_WOP(&updatestateAPI.UpdateState_WOP)

		backRepoData.UpdateStateAPIs = append(backRepoData.UpdateStateAPIs, &updatestateAPI)
	}

}
