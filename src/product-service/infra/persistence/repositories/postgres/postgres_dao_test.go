package postgres

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDAO_Write_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &DAO{
		DB: db,
	}

	query := "INSERT INTO users"
	mock.ExpectExec(query).WithArgs("test").WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := dao.Write(query, "test")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		t.Errorf("expected 1 row affected, got %d", rowsAffected)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDAO_Write_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &DAO{
		DB: db,
	}

	query := "INSERT INTO users"
	mock.ExpectExec(query).WithArgs("test").WillReturnError(fmt.Errorf("db error"))

	_, err = dao.Write(query, "test")
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDAO_Write_NilDB(t *testing.T) {
	dao := &DAO{DB: nil}
	_, err := dao.Write("SELECT 1")
	if err == nil {
		t.Errorf("expected error for nil DB")
	}
}

func TestDAO_Read_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &DAO{
		DB: db,
	}

	query := "SELECT id FROM users"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(query).WithArgs("test").WillReturnRows(rows)

	resultRows, err := dao.Read(query, "test")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	defer resultRows.Close()

	if !resultRows.Next() {
		t.Errorf("expected to read a row")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDAO_Read_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &DAO{
		DB: db,
	}

	query := "SELECT id FROM users"
	mock.ExpectQuery(query).WithArgs("test").WillReturnError(fmt.Errorf("db error"))

	_, err = dao.Read(query, "test")
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDAO_Read_NilDB(t *testing.T) {
	dao := &DAO{DB: nil}
	_, err := dao.Read("SELECT 1")
	if err == nil {
		t.Errorf("expected error for nil DB")
	}
}
