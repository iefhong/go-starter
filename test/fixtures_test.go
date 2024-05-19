package test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"iefhong/aw/go-starter/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	host            = os.Getenv("PSQL_HOST")
	port     int64  = 5432
	user            = os.Getenv("PSQL_USER")
	password string = os.Getenv("PSQL_PASS")
	dbname          = os.Getenv("PSQL_DBNAME")
)

func TestFixturesThroughSQLBoiler(t *testing.T) {

	fmt.Println("Connecting...")

	boil.DebugMode = true

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// trunc.
	models.Jets().DeleteAllP(context.Background(), db)
	models.PilotLanguages().DeleteAllP(context.Background(), db)
	models.Languages().DeleteAllP(context.Background(), db)
	models.Pilots().DeleteAllP(context.Background(), db)

	pilots, languages, pilotLanguages, jets := GetFixtures()
	t.Log(pilots)

	tx, err := db.BeginTx(context.TODO(), nil)

	if err != nil {
		t.Error("transaction fail")
	}

	for _, item := range pilots {
		item.InsertP(context.Background(), db, boil.Infer())
	}

	for _, item := range languages {
		item.InsertP(context.Background(), db, boil.Infer())
	}

	for _, item := range pilotLanguages {
		item.InsertP(context.Background(), db, boil.Infer())
	}

	for _, item := range jets {
		item.InsertP(context.Background(), db, boil.Infer())
	}

	// Rollback or commit
	err = tx.Commit()

	if err != nil {
		t.Error("transaction commit failed")
	}

}
