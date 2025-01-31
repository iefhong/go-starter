package pgtestpool

import (
	"context"
	"database/sql"
	"sync"
	"testing"
	"time"

	"github.com/friendsofgo/errors"
)

func testManagerFromEnv() *Manager {
	conf := DefaultManagerConfigFromEnv()
	conf.DatabasePrefix = "pgtestpool" // ensure we don't overlap with other pools running concurrently
	return NewManager(conf)
}

// test helpers should never return errors, but are passed the *testing.T instance and fail if needed. It seems to be recommended helper functions are moved to a testing.go file...
// https://medium.com/@povilasve/go-advanced-tips-tricks-a872503ac859
// https://about.sourcegraph.com/go/advanced-testing-in-go

func disconnectManager(t *testing.T, m *Manager) {

	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := m.Disconnect(ctx, true); err != nil {
		t.Logf("received error while disconnecting manager: %v", err)
	}
}

func initTemplateDB(wg *sync.WaitGroup, errs chan<- error, m *Manager) {
	defer wg.Done()

	template, err := m.InitializeTemplateDatabase(context.Background(), "hashinghash")
	if err != nil {
		errs <- err
		return
	}

	if template.Ready() {
		errs <- errors.New("template database is marked as ready")
		return
	}

	errs <- nil
}

func populateTemplateDB(t *testing.T, template *TemplateDatabase) {
	t.Helper()

	db, err := sql.Open("postgres", template.Config.ConnectionString())
	if err != nil {
		t.Fatalf("failed to open template database connection: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("failed to ping template database connection: %v", err)
	}

	if _, err := db.ExecContext(ctx, `
		CREATE EXTENSION "uuid-ossp";
		CREATE TABLE pilots (
			id uuid NOT NULL DEFAULT uuid_generate_v4(),
			"name" text NOT NULL,
			created_at timestamptz NOT NULL,
			updated_at timestamptz NULL,
			CONSTRAINT pilot_pkey PRIMARY KEY (id)
		);
		CREATE TABLE jets (
			id uuid NOT NULL DEFAULT uuid_generate_v4(),
			pilot_id uuid NOT NULL,
			age int4 NOT NULL,
			"name" text NOT NULL,
			color text NOT NULL,
			created_at timestamptz NOT NULL,
			updated_at timestamptz NULL,
			CONSTRAINT jet_pkey PRIMARY KEY (id)
		);
		ALTER TABLE jets ADD CONSTRAINT jet_pilots_fkey FOREIGN KEY (pilot_id) REFERENCES pilots(id);
	`); err != nil {
		t.Fatalf("failed to create tables in template database: %v", err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("failed to create transaction for test data in template database: %v", err)
	}

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO pilots (id, "name", created_at, updated_at) VALUES ('744a1a87-5ef7-4309-8814-0f1054751156', 'Mario', '2020-03-23 09:44:00.548', '2020-03-23 09:44:00.548');
		INSERT INTO pilots (id, "name", created_at, updated_at) VALUES ('20d9d155-2e95-49a2-8889-2ae975a8617e', 'Nick', '2020-03-23 09:44:00.548', '2020-03-23 09:44:00.548');
		INSERT INTO jets (id, pilot_id, age, "name", color, created_at, updated_at) VALUES ('67d9d0c7-34e5-48b0-9c7d-c6344995353c', '744a1a87-5ef7-4309-8814-0f1054751156', 26, 'F-14B', 'grey', '2020-03-23 09:44:00.000', '2020-03-23 09:44:00.000');
		INSERT INTO jets (id, pilot_id, age, "name", color, created_at, updated_at) VALUES ('facaf791-21b4-401a-bbac-67079ae4921f', '20d9d155-2e95-49a2-8889-2ae975a8617e', 27, 'F-14B', 'grey/red', '2020-03-23 09:44:00.000', '2020-03-23 09:44:00.000');
	`); err != nil {
		t.Fatalf("failed to insert test data into tables in template database: %v", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("failed to commit transaction with test data in template database: %v", err)
	}
}

func verifyTestDB(t *testing.T, test *TestDatabase) {
	t.Helper()

	db, err := sql.Open("postgres", test.Config.ConnectionString())
	if err != nil {
		t.Fatalf("failed to open test database connection: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("failed to ping test database connection: %v", err)
	}

	var pilotCount int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM pilots").Scan(&pilotCount); err != nil {
		t.Fatalf("failed to query pilot test data count: %v", err)
	}

	if pilotCount != 2 {
		t.Errorf("invalid pilot test data count, got %d, want 2", pilotCount)
	}

	var jetCount int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM jets").Scan(&jetCount); err != nil {
		t.Fatalf("failed to query jet test data count: %v", err)
	}

	if jetCount != 2 {
		t.Errorf("invalid jet test data count, got %d, want 2", jetCount)
	}
}

func getTestDB(wg *sync.WaitGroup, errs chan<- error, m *Manager) {
	defer wg.Done()

	db, err := m.GetTestDatabase(context.Background(), "hashinghash")
	if err != nil {
		errs <- err
		return
	}

	if !db.Ready() {
		errs <- errors.New("test database is marked as not ready")
		return
	}
	if !db.Dirty() {
		errs <- errors.New("test database is not marked as dirty")
	}

	errs <- nil
}
