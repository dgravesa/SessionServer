package data

import (
	"database/sql"

	"bitbucket.org/dangravesteam/WaterLogger/SessionServer/model"

	"bitbucket.org/waterlogger/dbcommon/dbserver"
)

// DBLayer is a data layer using a SQL backend.
type DBLayer struct {
	db *sql.DB
}

// NewDBLayer constructs a new data layer that uses a SQL database backend.
func NewDBLayer(cfgname string) (*DBLayer, error) {
	// create database or connect to existing
	db, err := dbserver.StartupDB(cfgname, "session_server")
	if err != nil {
		return nil, err
	}

	// create user sessions table
	_, err = db.Exec(dbCreateTableIfDoesNotExistQuery)

	return &DBLayer{db: db}, err
}

// AddSession inserts a new session into the session store.
func (l *DBLayer) AddSession(s model.Session) {
	l.db.Exec(dbInsertSessionQuery, s.UserID, s.KeyHash)
}

// RemoveSession removes the corresponding session from the store if it exists.
func (l *DBLayer) RemoveSession(s model.Session) {
	l.db.Exec(dbDeleteSessionQuery, s.UserID, s.KeyHash)
}

// IsValid returns true if the session matches a session in the store, otherwise false.
func (l *DBLayer) IsValid(s model.Session) bool {
	var sessionFound bool
	l.db.QueryRow(dbFindSessionQuery, s.UserID, s.KeyHash).Scan(&sessionFound)
	return sessionFound
}
