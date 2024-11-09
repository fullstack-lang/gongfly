// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongleaflet/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	circleDBs map[uint]*CircleDB

	nextIDCircleDB uint

	diviconDBs map[uint]*DivIconDB

	nextIDDivIconDB uint

	layergroupDBs map[uint]*LayerGroupDB

	nextIDLayerGroupDB uint

	layergroupuseDBs map[uint]*LayerGroupUseDB

	nextIDLayerGroupUseDB uint

	mapoptionsDBs map[uint]*MapOptionsDB

	nextIDMapOptionsDB uint

	markerDBs map[uint]*MarkerDB

	nextIDMarkerDB uint

	userclickDBs map[uint]*UserClickDB

	nextIDUserClickDB uint

	vlineDBs map[uint]*VLineDB

	nextIDVLineDB uint

	visualtrackDBs map[uint]*VisualTrackDB

	nextIDVisualTrackDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		circleDBs: make(map[uint]*CircleDB),

		diviconDBs: make(map[uint]*DivIconDB),

		layergroupDBs: make(map[uint]*LayerGroupDB),

		layergroupuseDBs: make(map[uint]*LayerGroupUseDB),

		mapoptionsDBs: make(map[uint]*MapOptionsDB),

		markerDBs: make(map[uint]*MarkerDB),

		userclickDBs: make(map[uint]*UserClickDB),

		vlineDBs: make(map[uint]*VLineDB),

		visualtrackDBs: make(map[uint]*VisualTrackDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *CircleDB:
		db.nextIDCircleDB++
		v.ID = db.nextIDCircleDB
		db.circleDBs[v.ID] = v
	case *DivIconDB:
		db.nextIDDivIconDB++
		v.ID = db.nextIDDivIconDB
		db.diviconDBs[v.ID] = v
	case *LayerGroupDB:
		db.nextIDLayerGroupDB++
		v.ID = db.nextIDLayerGroupDB
		db.layergroupDBs[v.ID] = v
	case *LayerGroupUseDB:
		db.nextIDLayerGroupUseDB++
		v.ID = db.nextIDLayerGroupUseDB
		db.layergroupuseDBs[v.ID] = v
	case *MapOptionsDB:
		db.nextIDMapOptionsDB++
		v.ID = db.nextIDMapOptionsDB
		db.mapoptionsDBs[v.ID] = v
	case *MarkerDB:
		db.nextIDMarkerDB++
		v.ID = db.nextIDMarkerDB
		db.markerDBs[v.ID] = v
	case *UserClickDB:
		db.nextIDUserClickDB++
		v.ID = db.nextIDUserClickDB
		db.userclickDBs[v.ID] = v
	case *VLineDB:
		db.nextIDVLineDB++
		v.ID = db.nextIDVLineDB
		db.vlineDBs[v.ID] = v
	case *VisualTrackDB:
		db.nextIDVisualTrackDB++
		v.ID = db.nextIDVisualTrackDB
		db.visualtrackDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CircleDB:
		delete(db.circleDBs, v.ID)
	case *DivIconDB:
		delete(db.diviconDBs, v.ID)
	case *LayerGroupDB:
		delete(db.layergroupDBs, v.ID)
	case *LayerGroupUseDB:
		delete(db.layergroupuseDBs, v.ID)
	case *MapOptionsDB:
		delete(db.mapoptionsDBs, v.ID)
	case *MarkerDB:
		delete(db.markerDBs, v.ID)
	case *UserClickDB:
		delete(db.userclickDBs, v.ID)
	case *VLineDB:
		delete(db.vlineDBs, v.ID)
	case *VisualTrackDB:
		delete(db.visualtrackDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CircleDB:
		db.circleDBs[v.ID] = v
		return db, nil
	case *DivIconDB:
		db.diviconDBs[v.ID] = v
		return db, nil
	case *LayerGroupDB:
		db.layergroupDBs[v.ID] = v
		return db, nil
	case *LayerGroupUseDB:
		db.layergroupuseDBs[v.ID] = v
		return db, nil
	case *MapOptionsDB:
		db.mapoptionsDBs[v.ID] = v
		return db, nil
	case *MarkerDB:
		db.markerDBs[v.ID] = v
		return db, nil
	case *UserClickDB:
		db.userclickDBs[v.ID] = v
		return db, nil
	case *VLineDB:
		db.vlineDBs[v.ID] = v
		return db, nil
	case *VisualTrackDB:
		db.visualtrackDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CircleDB:
		if existing, ok := db.circleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Circle github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *DivIconDB:
		if existing, ok := db.diviconDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DivIcon github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *LayerGroupDB:
		if existing, ok := db.layergroupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db LayerGroup github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *LayerGroupUseDB:
		if existing, ok := db.layergroupuseDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db LayerGroupUse github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *MapOptionsDB:
		if existing, ok := db.mapoptionsDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db MapOptions github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *MarkerDB:
		if existing, ok := db.markerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Marker github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *UserClickDB:
		if existing, ok := db.userclickDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db UserClick github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *VLineDB:
		if existing, ok := db.vlineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db VLine github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	case *VisualTrackDB:
		if existing, ok := db.visualtrackDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db VisualTrack github.com/fullstack-lang/gongleaflet/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]CircleDB:
		*ptr = make([]CircleDB, 0, len(db.circleDBs))
		for _, v := range db.circleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DivIconDB:
		*ptr = make([]DivIconDB, 0, len(db.diviconDBs))
		for _, v := range db.diviconDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LayerGroupDB:
		*ptr = make([]LayerGroupDB, 0, len(db.layergroupDBs))
		for _, v := range db.layergroupDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LayerGroupUseDB:
		*ptr = make([]LayerGroupUseDB, 0, len(db.layergroupuseDBs))
		for _, v := range db.layergroupuseDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MapOptionsDB:
		*ptr = make([]MapOptionsDB, 0, len(db.mapoptionsDBs))
		for _, v := range db.mapoptionsDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MarkerDB:
		*ptr = make([]MarkerDB, 0, len(db.markerDBs))
		for _, v := range db.markerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]UserClickDB:
		*ptr = make([]UserClickDB, 0, len(db.userclickDBs))
		for _, v := range db.userclickDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]VLineDB:
		*ptr = make([]VLineDB, 0, len(db.vlineDBs))
		for _, v := range db.vlineDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]VisualTrackDB:
		*ptr = make([]VisualTrackDB, 0, len(db.visualtrackDBs))
		for _, v := range db.visualtrackDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *CircleDB:
		tmp, ok := db.circleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Circle Unkown entry %d", i))
		}

		circleDB, _ := instanceDB.(*CircleDB)
		*circleDB = *tmp
		
	case *DivIconDB:
		tmp, ok := db.diviconDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DivIcon Unkown entry %d", i))
		}

		diviconDB, _ := instanceDB.(*DivIconDB)
		*diviconDB = *tmp
		
	case *LayerGroupDB:
		tmp, ok := db.layergroupDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First LayerGroup Unkown entry %d", i))
		}

		layergroupDB, _ := instanceDB.(*LayerGroupDB)
		*layergroupDB = *tmp
		
	case *LayerGroupUseDB:
		tmp, ok := db.layergroupuseDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First LayerGroupUse Unkown entry %d", i))
		}

		layergroupuseDB, _ := instanceDB.(*LayerGroupUseDB)
		*layergroupuseDB = *tmp
		
	case *MapOptionsDB:
		tmp, ok := db.mapoptionsDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First MapOptions Unkown entry %d", i))
		}

		mapoptionsDB, _ := instanceDB.(*MapOptionsDB)
		*mapoptionsDB = *tmp
		
	case *MarkerDB:
		tmp, ok := db.markerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Marker Unkown entry %d", i))
		}

		markerDB, _ := instanceDB.(*MarkerDB)
		*markerDB = *tmp
		
	case *UserClickDB:
		tmp, ok := db.userclickDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First UserClick Unkown entry %d", i))
		}

		userclickDB, _ := instanceDB.(*UserClickDB)
		*userclickDB = *tmp
		
	case *VLineDB:
		tmp, ok := db.vlineDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First VLine Unkown entry %d", i))
		}

		vlineDB, _ := instanceDB.(*VLineDB)
		*vlineDB = *tmp
		
	case *VisualTrackDB:
		tmp, ok := db.visualtrackDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First VisualTrack Unkown entry %d", i))
		}

		visualtrackDB, _ := instanceDB.(*VisualTrackDB)
		*visualtrackDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongleaflet/go, Unkown type")
	}
	
	return db, nil
}

