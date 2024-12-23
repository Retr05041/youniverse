package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestLoadDatabase_ValidFile(t *testing.T) {
	// Create a mock database connection
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create SQL mock: %v", err)
	}
	defer dbMock.Close()

	// Replace sql.Open with a mockable version for testing
	originalSQL := sqlOpen
	sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		if dataSourceName != "test_db.db" {
			return nil, errors.New("unexpected database name")
		}
		return dbMock, nil
	}
	defer func() { sqlOpen = originalSQL }()

	// Mock SQL expectations for table creation - TODO: Fix this.
	mock.ExpectExec(`
		CREATE TABLE IF NOT EXISTS players(
        player_id INTEGER PRIMARY KEY AUTOINCREMENT,
        player_name TEXT NOT NULL UNIQUE, 
        curr_room_index SMALLINT NOT NULL
		); 
		CREATE TABLE IF NOT EXISTS inventory(
        inventory_id INTEGER PRIMARY KEY AUTOINCREMENT,
        player_id INTEGER NOT NULL,
        item TEXT NOT NULL,
        FOREIGN KEY (player_id) REFERENCES players(player_id)
		);`).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call LoadDatabase
	db, err := LoadDatabase("test_db")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the database instance is valid
	if db.instance != dbMock {
		t.Fatalf("Expected mock database instance, got %v", db.instance)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("SQL expectations were not met: %v", err)
	}
}
