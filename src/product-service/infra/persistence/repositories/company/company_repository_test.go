package company

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	postgres "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres"
	queries "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres/queries"
)

func TestCompanyRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCompanyRepository(dao, queries.NewCompanyQueries())

	company := &domain.Company{Name: "Test Company"}

	mock.ExpectExec("INSERT INTO companies").
		WithArgs(company.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	compChan, errChan := repo.Save(company)

	select {
	case res := <-compChan:
		if res.Name != company.Name {
			t.Errorf("expected company name %s, got %s", company.Name, res.Name)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCompanyRepository_FindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCompanyRepository(dao, queries.NewCompanyQueries())

	id := "test-id"
	name := "Test Company"
	
	rows := sqlmock.NewRows([]string{"id", "name", "deleted_at"}).
		AddRow(id, name, nil)

	mock.ExpectQuery("SELECT id, name, deleted_at FROM companies").
		WithArgs(id).
		WillReturnRows(rows)

	compChan, errChan := repo.FindById(id)

	select {
	case res := <-compChan:
		if res == nil {
			t.Fatal("expected company, got nil")
		}
		if res.Id != id || res.Name != name {
			t.Errorf("expected id %s and name %s, got %s and %s", id, name, res.Id, res.Name)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCompanyRepository_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCompanyRepository(dao, queries.NewCompanyQueries())

	rows := sqlmock.NewRows([]string{"id", "name", "deleted_at"}).
		AddRow("1", "Comp 1", nil).
		AddRow("2", "Comp 2", nil)

	mock.ExpectQuery("SELECT id, name, deleted_at FROM companies").
		WillReturnRows(rows)

	compChan, errChan := repo.FindAll()

	select {
	case res := <-compChan:
		if len(res) != 2 {
			t.Errorf("expected 2 companies, got %d", len(res))
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCompanyRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCompanyRepository(dao, queries.NewCompanyQueries())

	company := &domain.Company{Id: "1", Name: "Updated Name"}

	mock.ExpectExec("UPDATE companies").
		WithArgs(company.Name, company.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Update(company)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCompanyRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCompanyRepository(dao, queries.NewCompanyQueries())

	id := "1"
	now := time.Now()

	mock.ExpectExec("UPDATE companies").
		WithArgs(now, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Delete(id, now)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}
