package sqltype

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func init() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}
	DB = db
}

func TestIntegrationDuration(t *testing.T) {
	_, err := DB.Exec("CREATE TABLE `durations` (took varchar(30))")
	if err != nil {
		t.Fatal(err)
	}
	nd := NullDuration{Valid: true, Duration: 5 * time.Second}
	_, err = DB.Exec("INSERT INTO `durations` (took) VALUES (?)", nd)
	if err != nil {
		t.Fatal(err)
	}

	type Task struct {
		Took NullDuration
	}
	var task Task

	row := DB.QueryRow("SELECT took FROM `durations` LIMIT 1")
	err = row.Scan(&task.Took)
	if err != nil {
		t.Fatal(err)
	}

	if want, have := nd.Valid, task.Took.Valid; want != have {
		t.Errorf("want Valid=%v, have %v", want, have)
	}
	if want, have := nd.Duration, task.Took.Duration; want != have {
		t.Errorf("want Duration=%v, have %v", want, have)
	}
}

func TestIntegrationTime(t *testing.T) {
	_, err := DB.Exec("CREATE TABLE `times` (`created` datetime)")
	if err != nil {
		t.Fatal(err)
	}
	nt := NullTime{Valid: true, Time: time.Date(2018, 8, 12, 14, 15, 39, 1234, time.UTC)}
	_, err = DB.Exec("INSERT INTO `times` (`created`) VALUES (?)", nt)
	if err != nil {
		t.Fatal(err)
	}

	type Task struct {
		Created NullTime
	}
	var task Task

	row := DB.QueryRow("SELECT `created` FROM `times` LIMIT 1")
	err = row.Scan(&task.Created)
	if err != nil {
		t.Fatal(err)
	}

	if want, have := nt.Valid, task.Created.Valid; want != have {
		t.Errorf("want Valid=%v, have %v", want, have)
	}
	if want, have := nt.Time, task.Created.Time; want != have {
		t.Errorf("want Time=%v, have %v", want, have)
	}
}
