package session

import (
	"fmt"
	"orm/log"
	"orm/schema"
	"reflect"
	"strings"
)

func (session *Session) Model(value interface{}) *Session {
	if session.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(session.refTable.Model) {
		session.refTable = schema.Parse(value, session.dialect)
	}
	return session
}

func (session *Session) RefTable() *schema.Schema {
	if session.refTable == nil {
		log.Error("Model is not set")
	}
	return session.refTable
}

func (session *Session) CreateTable() error {
	table := session.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := session.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

func (session *Session) DropTable() error {
	_, err := session.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", session.RefTable().Name)).Exec()
	return err
}

func (session *Session) HasTable() bool {
	sql, values := session.dialect.TableExistSQL(session.RefTable().Name)
	row := session.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == session.RefTable().Name
}
