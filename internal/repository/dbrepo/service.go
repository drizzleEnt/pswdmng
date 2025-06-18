package dbrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pswdmng/internal/domain"
	"pswdmng/internal/repository"
	"strings"
	"time"

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
		password text,
		deleted_at bigint
	)`)

	if err != nil {
		return err
	}

	return nil
}

// Add implements repository.Repository.
func (r *repo) Add(account string, login string, url string, psw string) error {
	db, err := db(account)
	if err != nil {
		return err
	}
	defer closeDB(db)

	query := `
	INSERT INTO passwords
	(login, url, password) 
	VALUES ($1, $2, $3)`

	_, err = db.Exec(query, login, url, psw)
	if err != nil {
		return err
	}
	return nil
}

// CheckExist implements repository.Repository.
func (r *repo) CheckExist() (bool, []domain.UserInfo, error) {
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

	query := `
	SELECT 
	password 
	FROM passwords 
	WHERE 
	(url = $1 AND login = $2) AND deleted_at IS NULL`

	var dbPswd string
	if err := db.QueryRow(query, url, login).Scan(&dbPswd); err != nil {
		return "", err
	}

	return dbPswd, nil
}

// List implements repository.Repository.
func (r *repo) List(login string) ([]domain.UserInfo, error) {
	db, err := db(login)
	if err != nil {
		return nil, err
	}
	defer closeDB(db)

	query := `
	SELECT 
	url, login 
	FROM passwords
	WHERE deleted_at IS NULL`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []domain.UserInfo

	for rows.Next() {
		var userInfo domain.UserInfo

		if err := rows.Scan(&userInfo.Url, &userInfo.Login); err != nil {
			return nil, err
		}

		entries = append(entries, userInfo)
	}

	return entries, nil
}

// Remove implements repository.Repository.
func (r *repo) Remove(account string, url string, login string) error {
	db, err := db(account)
	if err != nil {
		return err
	}
	defer db.Close()

	now := time.Now().Unix()
	query := `
	UPDATE passwords 
	SET 
	deleted_at = $1
	WHERE url = $2 AND login = $3`

	if _, err := db.Exec(query, now, url, login); err != nil {
		return err
	}

	return nil
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

func (r *repo) getExistFiles() ([]domain.UserInfo, error) {
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

	logins := make([]domain.UserInfo, 0, len(dirEntry))
	for _, entry := range dirEntry {
		var userInfo domain.UserInfo
		userInfo.Login = strings.Split(entry.Name(), "_")[0]
		logins = append(logins, userInfo)
	}

	return logins, nil
}
