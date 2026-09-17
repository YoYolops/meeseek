package keeper

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Keeper struct {
	DB *sql.DB
}

func SetupKeeper() (Keeper, error) {
	db, err := initDB()
	if err != nil {
		return Keeper{}, fmt.Errorf("InitDB_Error | %w", err)
	}
	return Keeper{
		db,
	}, nil
}

func initDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Failed to load .env file: %w", err)
	}

	dbPath := os.Getenv("KEEPER_DB_PATH")
	if dbPath == "" {
		return nil, fmt.Errorf("KEEPER_DB_PATH is not defined in .env file")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);`

	if _, err := db.Exec(query); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return db, nil
}

func (k Keeper) Retrieve() {
	fmt.Println("RETRIEVING")
}
