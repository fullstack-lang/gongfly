// generated by genORMTranslation.go
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/fullstack-lang/gongsim/go/models"

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoDummyAgent BackRepoDummyAgentStruct

	BackRepoEngine BackRepoEngineStruct

	BackRepoEvent BackRepoEventStruct

	BackRepoGongsimCommand BackRepoGongsimCommandStruct

	BackRepoGongsimStatus BackRepoGongsimStatusStruct

	CommitFromBackNb uint // this ng is updated at the BackRepo level but also at the BackRepo<GongStruct> level

	PushFromFrontNb uint // records increments from push from front
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromBackCallback != nil {
		models.Stage.OnInitCommitFromBackCallback.BeforeCommit(&models.Stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromFrontCallback != nil {
		models.Stage.OnInitCommitFromFrontCallback.BeforeCommit(&models.Stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) init(db *gorm.DB) {
	// insertion point for per struct back repo declarations
	backRepo.BackRepoDummyAgent.Init(db)
	backRepo.BackRepoEngine.Init(db)
	backRepo.BackRepoEvent.Init(db)
	backRepo.BackRepoGongsimCommand.Init(db)
	backRepo.BackRepoGongsimStatus.Init(db)

	models.Stage.BackRepo = backRepo
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoDummyAgent.CommitPhaseOne(stage)
	backRepo.BackRepoEngine.CommitPhaseOne(stage)
	backRepo.BackRepoEvent.CommitPhaseOne(stage)
	backRepo.BackRepoGongsimCommand.CommitPhaseOne(stage)
	backRepo.BackRepoGongsimStatus.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoDummyAgent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEngine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEvent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongsimCommand.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongsimStatus.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoDummyAgent.CheckoutPhaseOne()
	backRepo.BackRepoEngine.CheckoutPhaseOne()
	backRepo.BackRepoEvent.CheckoutPhaseOne()
	backRepo.BackRepoGongsimCommand.CheckoutPhaseOne()
	backRepo.BackRepoGongsimStatus.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoDummyAgent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEngine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEvent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongsimCommand.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongsimStatus.CheckoutPhaseTwo(backRepo)
}

var BackRepo BackRepoStruct

func GetLastCommitFromBackNb() uint {
	return BackRepo.GetLastCommitFromBackNb()
}

func GetLastPushFromFrontNb() uint {
	return BackRepo.GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoDummyAgent.Backup(dirPath)
	backRepo.BackRepoEngine.Backup(dirPath)
	backRepo.BackRepoEvent.Backup(dirPath)
	backRepo.BackRepoGongsimCommand.Backup(dirPath)
	backRepo.BackRepoGongsimStatus.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoDummyAgent.BackupXL(file)
	backRepo.BackRepoEngine.BackupXL(file)
	backRepo.BackRepoEvent.BackupXL(file)
	backRepo.BackRepoGongsimCommand.BackupXL(file)
	backRepo.BackRepoGongsimStatus.BackupXL(file)

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
	models.Stage.Commit()
	models.Stage.Reset()
	models.Stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoDummyAgent.RestorePhaseOne(dirPath)
	backRepo.BackRepoEngine.RestorePhaseOne(dirPath)
	backRepo.BackRepoEvent.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongsimCommand.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongsimStatus.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoDummyAgent.RestorePhaseTwo()
	backRepo.BackRepoEngine.RestorePhaseTwo()
	backRepo.BackRepoEvent.RestorePhaseTwo()
	backRepo.BackRepoGongsimCommand.RestorePhaseTwo()
	backRepo.BackRepoGongsimStatus.RestorePhaseTwo()

	models.Stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	models.Stage.Reset()

	// commit the cleaned stage
	models.Stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoDummyAgent.RestoreXLPhaseOne(file)
	backRepo.BackRepoEngine.RestoreXLPhaseOne(file)
	backRepo.BackRepoEvent.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongsimCommand.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongsimStatus.RestoreXLPhaseOne(file)

	// commit the restored stage
	models.Stage.Commit()
}
