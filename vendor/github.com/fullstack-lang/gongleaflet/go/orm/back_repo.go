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

	"github.com/fullstack-lang/gongleaflet/go/db"
	"github.com/fullstack-lang/gongleaflet/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gongleaflet/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoCircle BackRepoCircleStruct

	BackRepoDivIcon BackRepoDivIconStruct

	BackRepoLayerGroup BackRepoLayerGroupStruct

	BackRepoLayerGroupUse BackRepoLayerGroupUseStruct

	BackRepoMapOptions BackRepoMapOptionsStruct

	BackRepoMarker BackRepoMarkerStruct

	BackRepoUserClick BackRepoUserClickStruct

	BackRepoVLine BackRepoVLineStruct

	BackRepoVisualTrack BackRepoVisualTrackStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gongleaflet_go",
		&CircleDB{},
		&DivIconDB{},
		&LayerGroupDB{},
		&LayerGroupUseDB{},
		&MapOptionsDB{},
		&MarkerDB{},
		&UserClickDB{},
		&VLineDB{},
		&VisualTrackDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoCircle = BackRepoCircleStruct{
		Map_CircleDBID_CirclePtr: make(map[uint]*models.Circle, 0),
		Map_CircleDBID_CircleDB:  make(map[uint]*CircleDB, 0),
		Map_CirclePtr_CircleDBID: make(map[*models.Circle]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDivIcon = BackRepoDivIconStruct{
		Map_DivIconDBID_DivIconPtr: make(map[uint]*models.DivIcon, 0),
		Map_DivIconDBID_DivIconDB:  make(map[uint]*DivIconDB, 0),
		Map_DivIconPtr_DivIconDBID: make(map[*models.DivIcon]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLayerGroup = BackRepoLayerGroupStruct{
		Map_LayerGroupDBID_LayerGroupPtr: make(map[uint]*models.LayerGroup, 0),
		Map_LayerGroupDBID_LayerGroupDB:  make(map[uint]*LayerGroupDB, 0),
		Map_LayerGroupPtr_LayerGroupDBID: make(map[*models.LayerGroup]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLayerGroupUse = BackRepoLayerGroupUseStruct{
		Map_LayerGroupUseDBID_LayerGroupUsePtr: make(map[uint]*models.LayerGroupUse, 0),
		Map_LayerGroupUseDBID_LayerGroupUseDB:  make(map[uint]*LayerGroupUseDB, 0),
		Map_LayerGroupUsePtr_LayerGroupUseDBID: make(map[*models.LayerGroupUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMapOptions = BackRepoMapOptionsStruct{
		Map_MapOptionsDBID_MapOptionsPtr: make(map[uint]*models.MapOptions, 0),
		Map_MapOptionsDBID_MapOptionsDB:  make(map[uint]*MapOptionsDB, 0),
		Map_MapOptionsPtr_MapOptionsDBID: make(map[*models.MapOptions]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMarker = BackRepoMarkerStruct{
		Map_MarkerDBID_MarkerPtr: make(map[uint]*models.Marker, 0),
		Map_MarkerDBID_MarkerDB:  make(map[uint]*MarkerDB, 0),
		Map_MarkerPtr_MarkerDBID: make(map[*models.Marker]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUserClick = BackRepoUserClickStruct{
		Map_UserClickDBID_UserClickPtr: make(map[uint]*models.UserClick, 0),
		Map_UserClickDBID_UserClickDB:  make(map[uint]*UserClickDB, 0),
		Map_UserClickPtr_UserClickDBID: make(map[*models.UserClick]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVLine = BackRepoVLineStruct{
		Map_VLineDBID_VLinePtr: make(map[uint]*models.VLine, 0),
		Map_VLineDBID_VLineDB:  make(map[uint]*VLineDB, 0),
		Map_VLinePtr_VLineDBID: make(map[*models.VLine]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVisualTrack = BackRepoVisualTrackStruct{
		Map_VisualTrackDBID_VisualTrackPtr: make(map[uint]*models.VisualTrack, 0),
		Map_VisualTrackDBID_VisualTrackDB:  make(map[uint]*VisualTrackDB, 0),
		Map_VisualTrackPtr_VisualTrackDBID: make(map[*models.VisualTrack]uint, 0),

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
	backRepo.BackRepoCircle.CommitPhaseOne(stage)
	backRepo.BackRepoDivIcon.CommitPhaseOne(stage)
	backRepo.BackRepoLayerGroup.CommitPhaseOne(stage)
	backRepo.BackRepoLayerGroupUse.CommitPhaseOne(stage)
	backRepo.BackRepoMapOptions.CommitPhaseOne(stage)
	backRepo.BackRepoMarker.CommitPhaseOne(stage)
	backRepo.BackRepoUserClick.CommitPhaseOne(stage)
	backRepo.BackRepoVLine.CommitPhaseOne(stage)
	backRepo.BackRepoVisualTrack.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCircle.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDivIcon.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLayerGroup.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLayerGroupUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMapOptions.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMarker.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUserClick.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVLine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVisualTrack.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoCircle.CheckoutPhaseOne()
	backRepo.BackRepoDivIcon.CheckoutPhaseOne()
	backRepo.BackRepoLayerGroup.CheckoutPhaseOne()
	backRepo.BackRepoLayerGroupUse.CheckoutPhaseOne()
	backRepo.BackRepoMapOptions.CheckoutPhaseOne()
	backRepo.BackRepoMarker.CheckoutPhaseOne()
	backRepo.BackRepoUserClick.CheckoutPhaseOne()
	backRepo.BackRepoVLine.CheckoutPhaseOne()
	backRepo.BackRepoVisualTrack.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCircle.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDivIcon.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLayerGroup.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLayerGroupUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMapOptions.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMarker.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUserClick.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVLine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVisualTrack.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoCircle.Backup(dirPath)
	backRepo.BackRepoDivIcon.Backup(dirPath)
	backRepo.BackRepoLayerGroup.Backup(dirPath)
	backRepo.BackRepoLayerGroupUse.Backup(dirPath)
	backRepo.BackRepoMapOptions.Backup(dirPath)
	backRepo.BackRepoMarker.Backup(dirPath)
	backRepo.BackRepoUserClick.Backup(dirPath)
	backRepo.BackRepoVLine.Backup(dirPath)
	backRepo.BackRepoVisualTrack.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoCircle.BackupXL(file)
	backRepo.BackRepoDivIcon.BackupXL(file)
	backRepo.BackRepoLayerGroup.BackupXL(file)
	backRepo.BackRepoLayerGroupUse.BackupXL(file)
	backRepo.BackRepoMapOptions.BackupXL(file)
	backRepo.BackRepoMarker.BackupXL(file)
	backRepo.BackRepoUserClick.BackupXL(file)
	backRepo.BackRepoVLine.BackupXL(file)
	backRepo.BackRepoVisualTrack.BackupXL(file)

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
	backRepo.BackRepoCircle.RestorePhaseOne(dirPath)
	backRepo.BackRepoDivIcon.RestorePhaseOne(dirPath)
	backRepo.BackRepoLayerGroup.RestorePhaseOne(dirPath)
	backRepo.BackRepoLayerGroupUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoMapOptions.RestorePhaseOne(dirPath)
	backRepo.BackRepoMarker.RestorePhaseOne(dirPath)
	backRepo.BackRepoUserClick.RestorePhaseOne(dirPath)
	backRepo.BackRepoVLine.RestorePhaseOne(dirPath)
	backRepo.BackRepoVisualTrack.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCircle.RestorePhaseTwo()
	backRepo.BackRepoDivIcon.RestorePhaseTwo()
	backRepo.BackRepoLayerGroup.RestorePhaseTwo()
	backRepo.BackRepoLayerGroupUse.RestorePhaseTwo()
	backRepo.BackRepoMapOptions.RestorePhaseTwo()
	backRepo.BackRepoMarker.RestorePhaseTwo()
	backRepo.BackRepoUserClick.RestorePhaseTwo()
	backRepo.BackRepoVLine.RestorePhaseTwo()
	backRepo.BackRepoVisualTrack.RestorePhaseTwo()

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
	backRepo.BackRepoCircle.RestoreXLPhaseOne(file)
	backRepo.BackRepoDivIcon.RestoreXLPhaseOne(file)
	backRepo.BackRepoLayerGroup.RestoreXLPhaseOne(file)
	backRepo.BackRepoLayerGroupUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoMapOptions.RestoreXLPhaseOne(file)
	backRepo.BackRepoMarker.RestoreXLPhaseOne(file)
	backRepo.BackRepoUserClick.RestoreXLPhaseOne(file)
	backRepo.BackRepoVLine.RestoreXLPhaseOne(file)
	backRepo.BackRepoVisualTrack.RestoreXLPhaseOne(file)

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
