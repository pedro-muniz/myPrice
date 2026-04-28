package branch

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	postgres "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres"
	queries "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres/queries"
)

func TestBranchRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewBranchRepository(dao, queries.NewBranchQueries())

	branch := &domain.Branch{CompanyId: "comp-1", Name: "Test Branch"}

	mock.ExpectExec("INSERT INTO branches").
		WithArgs(branch.CompanyId, branch.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	branchChan, errChan := repo.Save(branch)

	select {
	case res := <-branchChan:
		if res.Name != branch.Name {
			t.Errorf("expected branch name %s, got %s", branch.Name, res.Name)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestBranchRepository_FindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewBranchRepository(dao, queries.NewBranchQueries())

	id := "test-id"
	compId := "comp-1"
	name := "Test Branch"
	
	rows := sqlmock.NewRows([]string{"id", "company_id", "name", "deleted_at"}).
		AddRow(id, compId, name, nil)

	mock.ExpectQuery("SELECT id, company_id, name, deleted_at FROM branches").
		WithArgs(compId, id).
		WillReturnRows(rows)

	branchChan, errChan := repo.FindById(compId, id)

	select {
	case res := <-branchChan:
		if res == nil {
			t.Fatal("expected branch, got nil")
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

func TestBranchRepository_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewBranchRepository(dao, queries.NewBranchQueries())

	compId := "comp-1"
	rows := sqlmock.NewRows([]string{"id", "company_id", "name", "deleted_at"}).
		AddRow("1", compId, "Branch 1", nil).
		AddRow("2", compId, "Branch 2", nil)

	mock.ExpectQuery("SELECT id, company_id, name, deleted_at FROM branches").
		WithArgs(compId).
		WillReturnRows(rows)

	branchChan, errChan := repo.FindAll(compId)

	select {
	case res := <-branchChan:
		if len(res) != 2 {
			t.Errorf("expected 2 branches, got %d", len(res))
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestBranchRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewBranchRepository(dao, queries.NewBranchQueries())

	branch := &domain.Branch{Id: "1", CompanyId: "comp-1", Name: "Updated Name"}

	mock.ExpectExec("UPDATE branches").
		WithArgs(branch.Name, branch.CompanyId, branch.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Update(branch)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestBranchRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewBranchRepository(dao, queries.NewBranchQueries())

	id := "1"
	compId := "comp-1"
	now := time.Now()

	mock.ExpectExec("UPDATE branches").
		WithArgs(now, compId, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Delete(compId, id, now)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}
