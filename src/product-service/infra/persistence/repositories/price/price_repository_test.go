package price

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	postgres "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres"
	queries "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres/queries"
)

func TestPriceRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewPriceRepository(dao, queries.NewPriceQueries())

	price := &domain.Price{
		CompanyId: "comp-1", BranchId: "branch-1",
		Gross: 100.0, Net: 80.0, Selling: 120.0, Recommended: 130.0,
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	mock.ExpectExec("INSERT INTO prices").
		WithArgs(price.CompanyId, price.BranchId, price.Gross, price.Net, price.Selling, price.Recommended, price.CreatedAt, price.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	priceChan, errChan := repo.Save(price)

	select {
	case res := <-priceChan:
		if res.Gross != price.Gross {
			t.Errorf("expected gross %f, got %f", price.Gross, res.Gross)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestPriceRepository_FindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewPriceRepository(dao, queries.NewPriceQueries())

	id := "test-id"
	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()
	
	rows := sqlmock.NewRows([]string{"id", "company_id", "branch_id", "gross", "net", "selling", "recommended", "created_at", "updated_at", "deleted_at"}).
		AddRow(id, compId, branchId, 100.0, 80.0, 120.0, 130.0, now, now, nil)

	mock.ExpectQuery("SELECT id, company_id, branch_id, gross, net, selling, recommended, created_at, updated_at, deleted_at FROM prices").
		WithArgs(compId, branchId, id).
		WillReturnRows(rows)

	priceChan, errChan := repo.FindById(compId, branchId, id)

	select {
	case res := <-priceChan:
		if res == nil {
			t.Fatal("expected price, got nil")
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

func TestPriceRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewPriceRepository(dao, queries.NewPriceQueries())

	price := &domain.Price{
		Id: "1", CompanyId: "comp-1", BranchId: "branch-1",
		Gross: 110.0, Net: 90.0, Selling: 130.0, Recommended: 140.0,
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec("UPDATE prices").
		WithArgs(price.Gross, price.Net, price.Selling, price.Recommended, price.UpdatedAt, price.CompanyId, price.BranchId, price.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Update(price)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestPriceRepository_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewPriceRepository(dao, queries.NewPriceQueries())

	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "company_id", "branch_id", "gross", "net", "selling", "recommended", "created_at", "updated_at", "deleted_at"}).
		AddRow("1", compId, branchId, 100.0, 80.0, 120.0, 130.0, now, now, nil).
		AddRow("2", compId, branchId, 200.0, 180.0, 220.0, 230.0, now, now, nil)

	mock.ExpectQuery("SELECT id, company_id, branch_id, gross, net, selling, recommended, created_at, updated_at, deleted_at FROM prices").
		WithArgs(compId, branchId).
		WillReturnRows(rows)

	priceChan, errChan := repo.FindAll(compId, branchId)

	select {
	case res := <-priceChan:
		if len(res) != 2 {
			t.Errorf("expected 2 prices, got %d", len(res))
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestPriceRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewPriceRepository(dao, queries.NewPriceQueries())

	id := "1"
	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()

	mock.ExpectExec("UPDATE prices").
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
