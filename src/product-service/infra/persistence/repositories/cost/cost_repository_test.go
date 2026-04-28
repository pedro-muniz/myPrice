package cost

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	postgres "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres"
	queries "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres/queries"
)

func TestCostRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCostRepository(dao, queries.NewCostQueries())

	cost := &domain.Cost{CompanyId: "comp-1", BranchId: "branch-1", Name: "Test Cost", Value: 10.5}

	mock.ExpectExec("INSERT INTO costs").
		WithArgs(cost.CompanyId, cost.BranchId, cost.Name, cost.Value).
		WillReturnResult(sqlmock.NewResult(1, 1))

	costChan, errChan := repo.Save(cost)

	select {
	case res := <-costChan:
		if res.Name != cost.Name {
			t.Errorf("expected cost name %s, got %s", cost.Name, res.Name)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCostRepository_FindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCostRepository(dao, queries.NewCostQueries())

	id := "test-id"
	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()
	
	rows := sqlmock.NewRows([]string{"id", "company_id", "branch_id", "name", "value", "created_at", "updated_at", "deleted_at"}).
		AddRow(id, compId, branchId, "Test Cost", 10.5, now, now, nil)

	mock.ExpectQuery("SELECT id, company_id, branch_id, name, value, created_at, updated_at, deleted_at FROM costs").
		WithArgs(compId, branchId, id).
		WillReturnRows(rows)

	costChan, errChan := repo.FindById(compId, branchId, id)

	select {
	case res := <-costChan:
		if res == nil {
			t.Fatal("expected cost, got nil")
		}
		if res.Id != id {
			t.Errorf("expected id %s, got %s", id, res.Id)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCostRepository_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCostRepository(dao, queries.NewCostQueries())

	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "company_id", "branch_id", "name", "value", "created_at", "updated_at", "deleted_at"}).
		AddRow("1", compId, branchId, "Cost 1", 10.0, now, now, nil).
		AddRow("2", compId, branchId, "Cost 2", 20.0, now, now, nil)

	mock.ExpectQuery("SELECT id, company_id, branch_id, name, value, created_at, updated_at, deleted_at FROM costs").
		WithArgs(compId, branchId).
		WillReturnRows(rows)

	costChan, errChan := repo.FindAll(compId, branchId)

	select {
	case res := <-costChan:
		if len(res) != 2 {
			t.Errorf("expected 2 costs, got %d", len(res))
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCostRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCostRepository(dao, queries.NewCostQueries())

	cost := &domain.Cost{Id: "1", CompanyId: "comp-1", BranchId: "branch-1", Name: "Updated", Value: 15.0, UpdatedAt: time.Now()}

	mock.ExpectExec("UPDATE costs").
		WithArgs(cost.Name, cost.Value, cost.UpdatedAt, cost.CompanyId, cost.BranchId, cost.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Update(cost)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestCostRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewCostRepository(dao, queries.NewCostQueries())

	id := "1"
	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()

	mock.ExpectExec("UPDATE costs").
		WithArgs(now, compId, branchId, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Delete(compId, branchId, id, now)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}
