// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CivilianAirportAPIs []*CivilianAirportAPI

	LinerAPIs []*LinerAPI

	MessageAPIs []*MessageAPI

	OpsLineAPIs []*OpsLineAPI

	RadarAPIs []*RadarAPI

	SatelliteAPIs []*SatelliteAPI

	ScenarioAPIs []*ScenarioAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, civilianairportDB := range backRepo.BackRepoCivilianAirport.Map_CivilianAirportDBID_CivilianAirportDB {

		var civilianairportAPI CivilianAirportAPI
		civilianairportAPI.ID = civilianairportDB.ID
		civilianairportAPI.CivilianAirportPointersEncoding = civilianairportDB.CivilianAirportPointersEncoding
		civilianairportDB.CopyBasicFieldsToCivilianAirport_WOP(&civilianairportAPI.CivilianAirport_WOP)

		backRepoData.CivilianAirportAPIs = append(backRepoData.CivilianAirportAPIs, &civilianairportAPI)
	}

	for _, linerDB := range backRepo.BackRepoLiner.Map_LinerDBID_LinerDB {

		var linerAPI LinerAPI
		linerAPI.ID = linerDB.ID
		linerAPI.LinerPointersEncoding = linerDB.LinerPointersEncoding
		linerDB.CopyBasicFieldsToLiner_WOP(&linerAPI.Liner_WOP)

		backRepoData.LinerAPIs = append(backRepoData.LinerAPIs, &linerAPI)
	}

	for _, messageDB := range backRepo.BackRepoMessage.Map_MessageDBID_MessageDB {

		var messageAPI MessageAPI
		messageAPI.ID = messageDB.ID
		messageAPI.MessagePointersEncoding = messageDB.MessagePointersEncoding
		messageDB.CopyBasicFieldsToMessage_WOP(&messageAPI.Message_WOP)

		backRepoData.MessageAPIs = append(backRepoData.MessageAPIs, &messageAPI)
	}

	for _, opslineDB := range backRepo.BackRepoOpsLine.Map_OpsLineDBID_OpsLineDB {

		var opslineAPI OpsLineAPI
		opslineAPI.ID = opslineDB.ID
		opslineAPI.OpsLinePointersEncoding = opslineDB.OpsLinePointersEncoding
		opslineDB.CopyBasicFieldsToOpsLine_WOP(&opslineAPI.OpsLine_WOP)

		backRepoData.OpsLineAPIs = append(backRepoData.OpsLineAPIs, &opslineAPI)
	}

	for _, radarDB := range backRepo.BackRepoRadar.Map_RadarDBID_RadarDB {

		var radarAPI RadarAPI
		radarAPI.ID = radarDB.ID
		radarAPI.RadarPointersEncoding = radarDB.RadarPointersEncoding
		radarDB.CopyBasicFieldsToRadar_WOP(&radarAPI.Radar_WOP)

		backRepoData.RadarAPIs = append(backRepoData.RadarAPIs, &radarAPI)
	}

	for _, satelliteDB := range backRepo.BackRepoSatellite.Map_SatelliteDBID_SatelliteDB {

		var satelliteAPI SatelliteAPI
		satelliteAPI.ID = satelliteDB.ID
		satelliteAPI.SatellitePointersEncoding = satelliteDB.SatellitePointersEncoding
		satelliteDB.CopyBasicFieldsToSatellite_WOP(&satelliteAPI.Satellite_WOP)

		backRepoData.SatelliteAPIs = append(backRepoData.SatelliteAPIs, &satelliteAPI)
	}

	for _, scenarioDB := range backRepo.BackRepoScenario.Map_ScenarioDBID_ScenarioDB {

		var scenarioAPI ScenarioAPI
		scenarioAPI.ID = scenarioDB.ID
		scenarioAPI.ScenarioPointersEncoding = scenarioDB.ScenarioPointersEncoding
		scenarioDB.CopyBasicFieldsToScenario_WOP(&scenarioAPI.Scenario_WOP)

		backRepoData.ScenarioAPIs = append(backRepoData.ScenarioAPIs, &scenarioAPI)
	}

}
