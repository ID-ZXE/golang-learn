package session

import (
	"database/sql"
	"orm/dialect"
	"orm/log"
	"orm/schema"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (session *Session) Clear() {
	session.sql.Reset()
	session.sqlVars = nil
}

func (session *Session) DB() *sql.DB {
	return session.db
}

func (session *Session) Exec() (result sql.Result, err error) {
	defer session.Clear()
	log.Info(session.sql.String(), session.sqlVars)
	if result, err = session.DB().Exec(session.sql.String(), session.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (session *Session) QueryRow() *sql.Row {
	defer session.Clear()
	log.Info(session.sql.String(), session.sqlVars)
	return session.DB().QueryRow(session.sql.String(), session.sqlVars...)
}

func (session *Session) QueryRows() (rows *sql.Rows, err error) {
	defer session.Clear()
	log.Info(session.sql.String(), session.sqlVars)
	if rows, err = session.DB().Query(session.sql.String(), session.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (session *Session) Raw(sql string, values ...interface{}) *Session {
	session.sql.WriteString(sql)
	session.sql.WriteString(" ")
	session.sqlVars = append(session.sqlVars, values...)
	return session
}
