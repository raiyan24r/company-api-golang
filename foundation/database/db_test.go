package mysqldb_test

import (
	mysqldb "company-api/foundation/database"
	"testing"
	"time"
)

func TestOpen_Success(t *testing.T) {
	cfg := mysqldb.Config{
		Host:            "localhost",
		Port:            3307,
		User:            "root",
		Password:        "rootpassword",
		Name:            "company_db",
		MaxIdleConns:    5,
		MaxOpenConns:    10,
		ConnMaxLifetime: time.Hour,
	}

	db, err := mysqldb.Open(cfg)
	if err != nil {
		t.Fatalf("Expected successful connection, got error: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Errorf("Expected successful ping, got error: %v", err)
	}
}

func TestOpen_QueryDatabase(t *testing.T) {
	cfg := mysqldb.Config{
		Host:            "localhost",
		Port:            3307,
		User:            "root",
		Password:        "rootpassword",
		Name:            "company_db",
		MaxIdleConns:    5,
		MaxOpenConns:    10,
		ConnMaxLifetime: time.Hour,
	}

	db, err := mysqldb.Open(cfg)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		t.Errorf("Failed to query database: %v", err)
	}

	if version == "" {
		t.Error("Expected non-empty MySQL version")
	}

	t.Logf("Connected to MySQL version: %s", version)
}

func TestOpen_InvalidCredentials(t *testing.T) {
	cfg := mysqldb.Config{
		Host:            "localhost",
		Port:            3307,
		User:            "invalid_user",
		Password:        "wrong_password",
		Name:            "company_db",
		MaxIdleConns:    5,
		MaxOpenConns:    10,
		ConnMaxLifetime: time.Hour,
	}

	db, err := mysqldb.Open(cfg)
	if err == nil {
		db.Close()
		t.Fatal("Expected connection to fail with invalid credentials")
	}

	t.Logf("Expected error received: %v", err)
}

func TestOpen_InvalidPort(t *testing.T) {
	cfg := mysqldb.Config{
		Host:            "localhost",
		Port:            9999,
		User:            "root",
		Password:        "rootpassword",
		Name:            "company_db",
		MaxIdleConns:    5,
		MaxOpenConns:    10,
		ConnMaxLifetime: time.Hour,
	}

	db, err := mysqldb.Open(cfg)
	if err == nil {
		db.Close()
		t.Fatal("Expected connection to fail with invalid port")
	}

	t.Logf("Expected error received: %v", err)
}

func TestOpen_ConnectionPoolSettings(t *testing.T) {
	cfg := mysqldb.Config{
		Host:            "localhost",
		Port:            3307,
		User:            "root",
		Password:        "rootpassword",
		Name:            "company_db",
		MaxIdleConns:    3,
		MaxOpenConns:    7,
		ConnMaxLifetime: 30 * time.Minute,
	}

	db, err := mysqldb.Open(cfg)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	stats := db.Stats()

	if stats.MaxOpenConnections != 7 {
		t.Errorf("Expected MaxOpenConnections=7, got %d", stats.MaxOpenConnections)
	}

	t.Logf("Connection pool stats: %+v", stats)
}
