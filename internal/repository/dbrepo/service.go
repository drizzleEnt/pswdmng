package dbrepo

import (
	"database/sql"
	"pswdmng/internal/repository"

	_ "github.com/mattn/go-sqlite3"
)

func New() repository.Repository {
	return &repo{}
}

type repo struct {
}

// Add implements repository.Repository.
func (r *repo) Add() {
	panic("unimplemented")
}

// CheckExist implements repository.Repository.
func (r *repo) CheckExist() (bool, error) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return false, err
	}

	return true, nil
}

// Get implements repository.Repository.
func (r *repo) Get() {
	panic("unimplemented")
}

// List implements repository.Repository.
func (r *repo) List() {
	panic("unimplemented")
}

// Remove implements repository.Repository.
func (r *repo) Remove() {
	panic("unimplemented")
}
