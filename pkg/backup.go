package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

// BackupPostgres backs up a PostgreSQL database to a file
type backupPostgres struct {
	Host     string
	User     string
	Password string
	Database string
}

// constructor function
func NewBackupPostgres(host, user, password, database string) backupPostgres {
	//validate input
	if host == "" {
		panic("missing host")

	}
	if user == "" {
		panic("missing user")

	}
	if password == "" {
		panic("missing password")

	}
	if database == "" {
		panic("missing database")
	}
	return backupPostgres{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,
	}
}

func (b *backupPostgres) Backup(backupFile string) error {
	// Build the pg_dump command
	cmd := exec.Command("pg_dump", "-h", b.Host, "-U", b.User, "-d", b.Database, "-w", "-f", backupFile)

	// Set environment variable for PostgreSQL password
	cmd.Env = append(os.Environ(), "PGPASSWORD="+b.Password)

	// Run the pg_dump command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run pg_dump: %w", err)
	}

	return nil
}
