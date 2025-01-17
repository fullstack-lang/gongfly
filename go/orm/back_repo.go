// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongfly/go/db"
	"github.com/fullstack-lang/gongfly/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gongfly/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoCivilianAirport BackRepoCivilianAirportStruct

	BackRepoLiner BackRepoLinerStruct

	BackRepoMessage BackRepoMessageStruct

	BackRepoOpsLine BackRepoOpsLineStruct

	BackRepoRadar BackRepoRadarStruct

	BackRepoSatellite BackRepoSatelliteStruct

	BackRepoScenario BackRepoScenarioStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gongfly_go",
		&CivilianAirportDB{},
		&LinerDB{},
		&MessageDB{},
		&OpsLineDB{},
		&RadarDB{},
		&SatelliteDB{},
		&ScenarioDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoCivilianAirport = BackRepoCivilianAirportStruct{
		Map_CivilianAirportDBID_CivilianAirportPtr: make(map[uint]*models.CivilianAirport, 0),
		Map_CivilianAirportDBID_CivilianAirportDB:  make(map[uint]*CivilianAirportDB, 0),
		Map_CivilianAirportPtr_CivilianAirportDBID: make(map[*models.CivilianAirport]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLiner = BackRepoLinerStruct{
		Map_LinerDBID_LinerPtr: make(map[uint]*models.Liner, 0),
		Map_LinerDBID_LinerDB:  make(map[uint]*LinerDB, 0),
		Map_LinerPtr_LinerDBID: make(map[*models.Liner]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMessage = BackRepoMessageStruct{
		Map_MessageDBID_MessagePtr: make(map[uint]*models.Message, 0),
		Map_MessageDBID_MessageDB:  make(map[uint]*MessageDB, 0),
		Map_MessagePtr_MessageDBID: make(map[*models.Message]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOpsLine = BackRepoOpsLineStruct{
		Map_OpsLineDBID_OpsLinePtr: make(map[uint]*models.OpsLine, 0),
		Map_OpsLineDBID_OpsLineDB:  make(map[uint]*OpsLineDB, 0),
		Map_OpsLinePtr_OpsLineDBID: make(map[*models.OpsLine]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRadar = BackRepoRadarStruct{
		Map_RadarDBID_RadarPtr: make(map[uint]*models.Radar, 0),
		Map_RadarDBID_RadarDB:  make(map[uint]*RadarDB, 0),
		Map_RadarPtr_RadarDBID: make(map[*models.Radar]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSatellite = BackRepoSatelliteStruct{
		Map_SatelliteDBID_SatellitePtr: make(map[uint]*models.Satellite, 0),
		Map_SatelliteDBID_SatelliteDB:  make(map[uint]*SatelliteDB, 0),
		Map_SatellitePtr_SatelliteDBID: make(map[*models.Satellite]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScenario = BackRepoScenarioStruct{
		Map_ScenarioDBID_ScenarioPtr: make(map[uint]*models.Scenario, 0),
		Map_ScenarioDBID_ScenarioDB:  make(map[uint]*ScenarioDB, 0),
		Map_ScenarioPtr_ScenarioDBID: make(map[*models.Scenario]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()

	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {

	// forbid read of back repo during commit
	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoCivilianAirport.CommitPhaseOne(stage)
	backRepo.BackRepoLiner.CommitPhaseOne(stage)
	backRepo.BackRepoMessage.CommitPhaseOne(stage)
	backRepo.BackRepoOpsLine.CommitPhaseOne(stage)
	backRepo.BackRepoRadar.CommitPhaseOne(stage)
	backRepo.BackRepoSatellite.CommitPhaseOne(stage)
	backRepo.BackRepoScenario.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCivilianAirport.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLiner.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMessage.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOpsLine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRadar.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSatellite.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScenario.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoCivilianAirport.CheckoutPhaseOne()
	backRepo.BackRepoLiner.CheckoutPhaseOne()
	backRepo.BackRepoMessage.CheckoutPhaseOne()
	backRepo.BackRepoOpsLine.CheckoutPhaseOne()
	backRepo.BackRepoRadar.CheckoutPhaseOne()
	backRepo.BackRepoSatellite.CheckoutPhaseOne()
	backRepo.BackRepoScenario.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCivilianAirport.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLiner.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMessage.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOpsLine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRadar.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSatellite.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScenario.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoCivilianAirport.Backup(dirPath)
	backRepo.BackRepoLiner.Backup(dirPath)
	backRepo.BackRepoMessage.Backup(dirPath)
	backRepo.BackRepoOpsLine.Backup(dirPath)
	backRepo.BackRepoRadar.Backup(dirPath)
	backRepo.BackRepoSatellite.Backup(dirPath)
	backRepo.BackRepoScenario.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoCivilianAirport.BackupXL(file)
	backRepo.BackRepoLiner.BackupXL(file)
	backRepo.BackRepoMessage.BackupXL(file)
	backRepo.BackRepoOpsLine.BackupXL(file)
	backRepo.BackRepoRadar.BackupXL(file)
	backRepo.BackRepoSatellite.BackupXL(file)
	backRepo.BackRepoScenario.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCivilianAirport.RestorePhaseOne(dirPath)
	backRepo.BackRepoLiner.RestorePhaseOne(dirPath)
	backRepo.BackRepoMessage.RestorePhaseOne(dirPath)
	backRepo.BackRepoOpsLine.RestorePhaseOne(dirPath)
	backRepo.BackRepoRadar.RestorePhaseOne(dirPath)
	backRepo.BackRepoSatellite.RestorePhaseOne(dirPath)
	backRepo.BackRepoScenario.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCivilianAirport.RestorePhaseTwo()
	backRepo.BackRepoLiner.RestorePhaseTwo()
	backRepo.BackRepoMessage.RestorePhaseTwo()
	backRepo.BackRepoOpsLine.RestorePhaseTwo()
	backRepo.BackRepoRadar.RestorePhaseTwo()
	backRepo.BackRepoSatellite.RestorePhaseTwo()
	backRepo.BackRepoScenario.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCivilianAirport.RestoreXLPhaseOne(file)
	backRepo.BackRepoLiner.RestoreXLPhaseOne(file)
	backRepo.BackRepoMessage.RestoreXLPhaseOne(file)
	backRepo.BackRepoOpsLine.RestoreXLPhaseOne(file)
	backRepo.BackRepoRadar.RestoreXLPhaseOne(file)
	backRepo.BackRepoSatellite.RestoreXLPhaseOne(file)
	backRepo.BackRepoScenario.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.subscribersRwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.subscribersRwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.subscribersRwMutex.Lock()
	defer backRepoStruct.subscribersRwMutex.Unlock()
	for i, subscriber := range backRepoStruct.subscribers {
		if subscriber == ch {
			backRepoStruct.subscribers =
				append(backRepoStruct.subscribers[:i],
					backRepoStruct.subscribers[i+1:]...)
			close(ch) // Close the channel to signal completion
			break
		}
	}
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.subscribersRwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.subscribersRwMutex.RUnlock()

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}
