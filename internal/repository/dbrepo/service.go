package dbrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pswdmng/internal/repository"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func New() repository.Repository {
	return &repo{}
}

type repo struct {
}

// CreateFile implements repository.Repository.
func (r *repo) CreateFile(login string) error {
	db, err := db(login)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE passwords(
		url text,
		login text,
		password text
	)`)

	if err != nil {
		return err
	}

	return nil
}

// Add implements repository.Repository.
func (r *repo) Add() {
	panic("unimplemented")
}

// CheckExist implements repository.Repository.
func (r *repo) CheckExist() (bool, []string, error) {
	logins, err := r.getExistFiles()
	if err != nil {
		return false, nil, err
	}

	if len(logins) != 0 {
		return true, logins, nil
	}

	return false, nil, nil
}

// Get implements repository.Repository.
func (r *repo) Get() {
	panic("unimplemented")
}

// List implements repository.Repository.
func (r *repo) List(login string) ([]string, error) {
	db, err := db(login)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT url, login FROM passwords`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := make([]string, 0, 0)
	for rows.Next() {
		var login, url string
		if err := rows.Scan(&url, &login); err != nil {
			return nil, err
		}

		entries = append(entries, strings.Join([]string{url, login}, " "))

	}

	return entries, nil
}

// Remove implements repository.Repository.
func (r *repo) Remove() {
	panic("unimplemented")
}

func db(login string) (*sql.DB, error) {
	storeDir, err := getStoreDirPath()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", storeDir+"/"+fmt.Sprintf("%s_data.db", login))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getStoreDirPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	storeDir := filepath.Join(homeDir, ".pswdmng")
	return storeDir, nil
}

func closeDB(db *sql.DB) {
	db.Close()
}

func (r *repo) getExistFiles() ([]string, error) {
	storeDir, err := getStoreDirPath()
	if err != nil {
		return nil, err
	}

	dirEntry, err := os.ReadDir(storeDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.Mkdir(storeDir, 0700); err != nil {
				return nil, err
			}
			dirEntry, err = os.ReadDir(storeDir)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	logins := make([]string, 0, len(dirEntry))
	for _, entry := range dirEntry {
		foundedLogins := strings.Split(entry.Name(), "_")
		logins = append(logins, foundedLogins[0])
	}

	return logins, nil
}
