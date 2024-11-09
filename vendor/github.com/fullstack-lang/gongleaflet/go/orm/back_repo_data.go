// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CircleAPIs []*CircleAPI

	DivIconAPIs []*DivIconAPI

	LayerGroupAPIs []*LayerGroupAPI

	LayerGroupUseAPIs []*LayerGroupUseAPI

	MapOptionsAPIs []*MapOptionsAPI

	MarkerAPIs []*MarkerAPI

	UserClickAPIs []*UserClickAPI

	VLineAPIs []*VLineAPI

	VisualTrackAPIs []*VisualTrackAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, circleDB := range backRepo.BackRepoCircle.Map_CircleDBID_CircleDB {

		var circleAPI CircleAPI
		circleAPI.ID = circleDB.ID
		circleAPI.CirclePointersEncoding = circleDB.CirclePointersEncoding
		circleDB.CopyBasicFieldsToCircle_WOP(&circleAPI.Circle_WOP)

		backRepoData.CircleAPIs = append(backRepoData.CircleAPIs, &circleAPI)
	}

	for _, diviconDB := range backRepo.BackRepoDivIcon.Map_DivIconDBID_DivIconDB {

		var diviconAPI DivIconAPI
		diviconAPI.ID = diviconDB.ID
		diviconAPI.DivIconPointersEncoding = diviconDB.DivIconPointersEncoding
		diviconDB.CopyBasicFieldsToDivIcon_WOP(&diviconAPI.DivIcon_WOP)

		backRepoData.DivIconAPIs = append(backRepoData.DivIconAPIs, &diviconAPI)
	}

	for _, layergroupDB := range backRepo.BackRepoLayerGroup.Map_LayerGroupDBID_LayerGroupDB {

		var layergroupAPI LayerGroupAPI
		layergroupAPI.ID = layergroupDB.ID
		layergroupAPI.LayerGroupPointersEncoding = layergroupDB.LayerGroupPointersEncoding
		layergroupDB.CopyBasicFieldsToLayerGroup_WOP(&layergroupAPI.LayerGroup_WOP)

		backRepoData.LayerGroupAPIs = append(backRepoData.LayerGroupAPIs, &layergroupAPI)
	}

	for _, layergroupuseDB := range backRepo.BackRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB {

		var layergroupuseAPI LayerGroupUseAPI
		layergroupuseAPI.ID = layergroupuseDB.ID
		layergroupuseAPI.LayerGroupUsePointersEncoding = layergroupuseDB.LayerGroupUsePointersEncoding
		layergroupuseDB.CopyBasicFieldsToLayerGroupUse_WOP(&layergroupuseAPI.LayerGroupUse_WOP)

		backRepoData.LayerGroupUseAPIs = append(backRepoData.LayerGroupUseAPIs, &layergroupuseAPI)
	}

	for _, mapoptionsDB := range backRepo.BackRepoMapOptions.Map_MapOptionsDBID_MapOptionsDB {

		var mapoptionsAPI MapOptionsAPI
		mapoptionsAPI.ID = mapoptionsDB.ID
		mapoptionsAPI.MapOptionsPointersEncoding = mapoptionsDB.MapOptionsPointersEncoding
		mapoptionsDB.CopyBasicFieldsToMapOptions_WOP(&mapoptionsAPI.MapOptions_WOP)

		backRepoData.MapOptionsAPIs = append(backRepoData.MapOptionsAPIs, &mapoptionsAPI)
	}

	for _, markerDB := range backRepo.BackRepoMarker.Map_MarkerDBID_MarkerDB {

		var markerAPI MarkerAPI
		markerAPI.ID = markerDB.ID
		markerAPI.MarkerPointersEncoding = markerDB.MarkerPointersEncoding
		markerDB.CopyBasicFieldsToMarker_WOP(&markerAPI.Marker_WOP)

		backRepoData.MarkerAPIs = append(backRepoData.MarkerAPIs, &markerAPI)
	}

	for _, userclickDB := range backRepo.BackRepoUserClick.Map_UserClickDBID_UserClickDB {

		var userclickAPI UserClickAPI
		userclickAPI.ID = userclickDB.ID
		userclickAPI.UserClickPointersEncoding = userclickDB.UserClickPointersEncoding
		userclickDB.CopyBasicFieldsToUserClick_WOP(&userclickAPI.UserClick_WOP)

		backRepoData.UserClickAPIs = append(backRepoData.UserClickAPIs, &userclickAPI)
	}

	for _, vlineDB := range backRepo.BackRepoVLine.Map_VLineDBID_VLineDB {

		var vlineAPI VLineAPI
		vlineAPI.ID = vlineDB.ID
		vlineAPI.VLinePointersEncoding = vlineDB.VLinePointersEncoding
		vlineDB.CopyBasicFieldsToVLine_WOP(&vlineAPI.VLine_WOP)

		backRepoData.VLineAPIs = append(backRepoData.VLineAPIs, &vlineAPI)
	}

	for _, visualtrackDB := range backRepo.BackRepoVisualTrack.Map_VisualTrackDBID_VisualTrackDB {

		var visualtrackAPI VisualTrackAPI
		visualtrackAPI.ID = visualtrackDB.ID
		visualtrackAPI.VisualTrackPointersEncoding = visualtrackDB.VisualTrackPointersEncoding
		visualtrackDB.CopyBasicFieldsToVisualTrack_WOP(&visualtrackAPI.VisualTrack_WOP)

		backRepoData.VisualTrackAPIs = append(backRepoData.VisualTrackAPIs, &visualtrackAPI)
	}

}
