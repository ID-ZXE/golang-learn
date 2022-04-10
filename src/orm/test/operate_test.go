package test

import (
	"database/sql"
	"log"
	"orm"
	"orm/dialect"
	"orm/schema"
	"orm/session"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	url        = "root:root@tcp(127.0.0.1:3306)/foobar?charset=utf8"
)

var TestDial, _ = dialect.GetDialect("mysql")

func Test1(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/foobar?charset=utf8")
	if err != nil {
		panic(err)
	}
	// 检查是否连接成功数据库
	// db.Ping()

	result, err := db.Exec("INSERT INTO home_page(`author_id`,`title`, `content`) values (?,?,?)", 100, "golang title", "golang content")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println("affected", affected)
	}
	id, _ := result.LastInsertId()
	log.Println("LastInsertId", id)

	row := db.QueryRow("SELECT title FROM home_page LIMIT 10")
	var title string
	if err := row.Scan(&title); err == nil {
		log.Println(title)
	}
}

/********************************/

type User struct {
	Name string `orm:"PRIMARY KEY"`
	Age  int
}

func TestParse(t *testing.T) {
	schemaObj := schema.Parse(&User{}, TestDial)
	if schemaObj.Name != "User" || len(schemaObj.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schemaObj.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}

func TestSchemaRecordValues(t *testing.T) {
	schemaObj := schema.Parse(&User{}, TestDial)
	values := schemaObj.RecordValues(&User{"Tom", 18})

	name := values[0].(string)
	age := values[1].(int)

	if name != "Tom" || age != 18 {
		t.Fatal("failed to get values")
	}
}

/********************************/

type UserTest struct {
	Name string `orm:"PRIMARY KEY"`
	Age  int
}

func (u *UserTest) TableName() string {
	return "ns_user_test"
}

func TestSchemaTableName(t *testing.T) {
	schemaObj := schema.Parse(&UserTest{}, TestDial)
	if schemaObj.Name != "ns_user_test" || len(schemaObj.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
}

/********************************/

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}

func OpenDB(t *testing.T) *orm.Engine {
	t.Helper()
	engine, err := orm.NewEngine(driverName, url)
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

/********************************/

var (
	TestDB *sql.DB
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open(driverName, url)
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *session.Session {
	return session.New(TestDB, TestDial)
}

func TestSessionExec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2, but got", count)
	}
}

func TestRowExec(t *testing.T) {
	session := NewSession()
	_, _ = session.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = session.Raw("CREATE TABLE User(Name text);").Exec()
}

func TestQueryRows(t *testing.T) {
	session := NewSession()
	row := session.Raw("SELECT count(*) FROM User").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err)
	}
}

/********************************/

func TestSessionCreateTable(t *testing.T) {
	model := NewSession().Model(&User{})
	_ = model.DropTable()
	_ = model.CreateTable()
	if !model.HasTable() {
		t.Fatal("Failed to create table User")
	}
}

func TestSessionModel(t *testing.T) {
	model := NewSession().Model(&User{})
	table := model.RefTable()
	model.Model(&session.Session{})
	if table.Name != "User" || model.RefTable().Name != "Session" {
		t.Fatal("Failed to change model")
	}
}
