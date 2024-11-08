// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongsim/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	dummyagentDBs map[uint]*DummyAgentDB

	nextIDDummyAgentDB uint

	engineDBs map[uint]*EngineDB

	nextIDEngineDB uint

	eventDBs map[uint]*EventDB

	nextIDEventDB uint

	gongsimcommandDBs map[uint]*GongsimCommandDB

	nextIDGongsimCommandDB uint

	gongsimstatusDBs map[uint]*GongsimStatusDB

	nextIDGongsimStatusDB uint

	updatestateDBs map[uint]*UpdateStateDB

	nextIDUpdateStateDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		dummyagentDBs: make(map[uint]*DummyAgentDB),

		engineDBs: make(map[uint]*EngineDB),

		eventDBs: make(map[uint]*EventDB),

		gongsimcommandDBs: make(map[uint]*GongsimCommandDB),

		gongsimstatusDBs: make(map[uint]*GongsimStatusDB),

		updatestateDBs: make(map[uint]*UpdateStateDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *DummyAgentDB:
		db.nextIDDummyAgentDB++
		v.ID = db.nextIDDummyAgentDB
		db.dummyagentDBs[v.ID] = v
	case *EngineDB:
		db.nextIDEngineDB++
		v.ID = db.nextIDEngineDB
		db.engineDBs[v.ID] = v
	case *EventDB:
		db.nextIDEventDB++
		v.ID = db.nextIDEventDB
		db.eventDBs[v.ID] = v
	case *GongsimCommandDB:
		db.nextIDGongsimCommandDB++
		v.ID = db.nextIDGongsimCommandDB
		db.gongsimcommandDBs[v.ID] = v
	case *GongsimStatusDB:
		db.nextIDGongsimStatusDB++
		v.ID = db.nextIDGongsimStatusDB
		db.gongsimstatusDBs[v.ID] = v
	case *UpdateStateDB:
		db.nextIDUpdateStateDB++
		v.ID = db.nextIDUpdateStateDB
		db.updatestateDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DummyAgentDB:
		delete(db.dummyagentDBs, v.ID)
	case *EngineDB:
		delete(db.engineDBs, v.ID)
	case *EventDB:
		delete(db.eventDBs, v.ID)
	case *GongsimCommandDB:
		delete(db.gongsimcommandDBs, v.ID)
	case *GongsimStatusDB:
		delete(db.gongsimstatusDBs, v.ID)
	case *UpdateStateDB:
		delete(db.updatestateDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DummyAgentDB:
		db.dummyagentDBs[v.ID] = v
		return db, nil
	case *EngineDB:
		db.engineDBs[v.ID] = v
		return db, nil
	case *EventDB:
		db.eventDBs[v.ID] = v
		return db, nil
	case *GongsimCommandDB:
		db.gongsimcommandDBs[v.ID] = v
		return db, nil
	case *GongsimStatusDB:
		db.gongsimstatusDBs[v.ID] = v
		return db, nil
	case *UpdateStateDB:
		db.updatestateDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DummyAgentDB:
		if existing, ok := db.dummyagentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DummyAgent github.com/fullstack-lang/gongsim/go, record not found")
		}
	case *EngineDB:
		if existing, ok := db.engineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Engine github.com/fullstack-lang/gongsim/go, record not found")
		}
	case *EventDB:
		if existing, ok := db.eventDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Event github.com/fullstack-lang/gongsim/go, record not found")
		}
	case *GongsimCommandDB:
		if existing, ok := db.gongsimcommandDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongsimCommand github.com/fullstack-lang/gongsim/go, record not found")
		}
	case *GongsimStatusDB:
		if existing, ok := db.gongsimstatusDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db GongsimStatus github.com/fullstack-lang/gongsim/go, record not found")
		}
	case *UpdateStateDB:
		if existing, ok := db.updatestateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db UpdateState github.com/fullstack-lang/gongsim/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]DummyAgentDB:
		*ptr = make([]DummyAgentDB, 0, len(db.dummyagentDBs))
		for _, v := range db.dummyagentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]EngineDB:
		*ptr = make([]EngineDB, 0, len(db.engineDBs))
		for _, v := range db.engineDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]EventDB:
		*ptr = make([]EventDB, 0, len(db.eventDBs))
		for _, v := range db.eventDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongsimCommandDB:
		*ptr = make([]GongsimCommandDB, 0, len(db.gongsimcommandDBs))
		for _, v := range db.gongsimcommandDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]GongsimStatusDB:
		*ptr = make([]GongsimStatusDB, 0, len(db.gongsimstatusDBs))
		for _, v := range db.gongsimstatusDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]UpdateStateDB:
		*ptr = make([]UpdateStateDB, 0, len(db.updatestateDBs))
		for _, v := range db.updatestateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongsim/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *DummyAgentDB:
		tmp, ok := db.dummyagentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DummyAgent Unkown entry %d", i))
		}

		dummyagentDB, _ := instanceDB.(*DummyAgentDB)
		*dummyagentDB = *tmp
		
	case *EngineDB:
		tmp, ok := db.engineDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Engine Unkown entry %d", i))
		}

		engineDB, _ := instanceDB.(*EngineDB)
		*engineDB = *tmp
		
	case *EventDB:
		tmp, ok := db.eventDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Event Unkown entry %d", i))
		}

		eventDB, _ := instanceDB.(*EventDB)
		*eventDB = *tmp
		
	case *GongsimCommandDB:
		tmp, ok := db.gongsimcommandDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongsimCommand Unkown entry %d", i))
		}

		gongsimcommandDB, _ := instanceDB.(*GongsimCommandDB)
		*gongsimcommandDB = *tmp
		
	case *GongsimStatusDB:
		tmp, ok := db.gongsimstatusDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First GongsimStatus Unkown entry %d", i))
		}

		gongsimstatusDB, _ := instanceDB.(*GongsimStatusDB)
		*gongsimstatusDB = *tmp
		
	case *UpdateStateDB:
		tmp, ok := db.updatestateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First UpdateState Unkown entry %d", i))
		}

		updatestateDB, _ := instanceDB.(*UpdateStateDB)
		*updatestateDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongsim/go, Unkown type")
	}
	
	return db, nil
}

