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
func (r *repo) Add(account string, login string, url string) error {
	db, err := db(account)
	if err != nil {
		return err
	}
	defer closeDB(db)

	_, err = db.Exec(`INSERT INTO passwords (login, url, password) VALUES ($1, $2, $3)`, login, url, "123")
	if err != nil {
		return err
	}

	return nil
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
func (r *repo) Get(account string, url string, login string) (string, error) {
	db, err := db(account)
	if err != nil {
		return "", err
	}
	defer closeDB(db)

	var dbPswd string
	err = db.QueryRow(`SELECT password FROM passwords WHERE url = $1 AND login = $2`, url, login).Scan(&dbPswd)
	if err != nil {
		return "", err
	}

	return dbPswd, nil
}

// List implements repository.Repository.
func (r *repo) List(login string) ([][]string, error) {
	db, err := db(login)
	if err != nil {
		return nil, err
	}
	defer closeDB(db)

	rows, err := db.Query(`SELECT url, login FROM passwords`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := make([][]string, 0, 0)
	for rows.Next() {
		dbRows := make([]string, 0, 2)
		var login, url string
		if err := rows.Scan(&url, &login); err != nil {
			return nil, err
		}

		dbRows = append(dbRows, login)
		dbRows = append(dbRows, url)
		entries = append(entries, dbRows)
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
