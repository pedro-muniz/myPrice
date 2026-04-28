package product

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	postgres "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres"
	queries "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories/postgres/queries"
)

func TestProductRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewProductRepository(dao, queries.NewProductQueries())

	product := &domain.Product{
		CompanyId: "comp-1", BranchId: "branch-1", BarCode: "123", Name: "Prod",
		Description: "Desc", Discount: 0, Price: domain.Price{Id: "price-1"},
		ProfitPercent: 10, Origin: domain.National, CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	mock.ExpectExec("INSERT INTO products").
		WithArgs(product.CompanyId, product.BranchId, product.BarCode, product.Name, product.Description,
			product.Discount, product.Price.Id, product.ProfitPercent, int(product.Origin), product.CreatedAt, product.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	prodChan, errChan := repo.Save(product)

	select {
	case res := <-prodChan:
		if res.Name != product.Name {
			t.Errorf("expected name %s, got %s", product.Name, res.Name)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestProductRepository_FindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewProductRepository(dao, queries.NewProductQueries())

	id := "prod-1"
	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()
	
	rows := sqlmock.NewRows([]string{
		"pd.id", "pd.company_id", "pd.branch_id", "pd.barcode", "pd.name", "pd.description", "pd.discount",
		"pd.profit_percent", "pd.origin", "pd.created_at", "pd.updated_at", "pd.deleted_at",
		"pr.id", "pr.gross", "pr.net", "pr.selling", "pr.recommended", "pr.created_at", "pr.updated_at", "pr.deleted_at",
	}).AddRow(
		id, compId, branchId, "123", "Prod", "Desc", 0.0, 10.0, 0, now, now, nil,
		"price-1", 100.0, 80.0, 120.0, 130.0, now, now, nil,
	)

	mock.ExpectQuery("SELECT pd.id, pd.company_id, pd.branch_id, pd.barcode, pd.name, pd.description, pd.discount, pd.profit_percent, pd.origin, pd.created_at, pd.updated_at, pd.deleted_at, pr.id, pr.gross, pr.net, pr.selling, pr.recommended, pr.created_at, pr.updated_at, pr.deleted_at FROM products pd JOIN prices pr ON pd.price_id = pr.id").
		WithArgs(compId, branchId, id).
		WillReturnRows(rows)

	prodChan, errChan := repo.FindById(compId, branchId, id)

	select {
	case res := <-prodChan:
		if res == nil {
			t.Fatal("expected product, got nil")
		}
		if res.Id != id || res.Price.Id != "price-1" {
			t.Errorf("expected product id %s and price id price-1, got %s and %s", id, res.Id, res.Price.Id)
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestProductRepository_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewProductRepository(dao, queries.NewProductQueries())

	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"pd.id", "pd.company_id", "pd.branch_id", "pd.barcode", "pd.name", "pd.description", "pd.discount",
		"pd.profit_percent", "pd.origin", "pd.created_at", "pd.updated_at", "pd.deleted_at",
		"pr.id", "pr.gross", "pr.net", "pr.selling", "pr.recommended", "pr.created_at", "pr.updated_at", "pr.deleted_at",
	}).AddRow(
		"prod-1", compId, branchId, "123", "Prod 1", "Desc 1", 0.0, 10.0, 0, now, now, nil,
		"price-1", 100.0, 80.0, 120.0, 130.0, now, now, nil,
	).AddRow(
		"prod-2", compId, branchId, "456", "Prod 2", "Desc 2", 5.0, 15.0, 1, now, now, nil,
		"price-2", 200.0, 180.0, 220.0, 230.0, now, now, nil,
	)

	mock.ExpectQuery("SELECT pd.id, pd.company_id, pd.branch_id, pd.barcode, pd.name, pd.description, pd.discount, pd.profit_percent, pd.origin, pd.created_at, pd.updated_at, pd.deleted_at, pr.id, pr.gross, pr.net, pr.selling, pr.recommended, pr.created_at, pr.updated_at, pr.deleted_at FROM products pd JOIN prices pr ON pd.price_id = pr.id").
		WithArgs(compId, branchId).
		WillReturnRows(rows)

	prodChan, errChan := repo.FindAll(compId, branchId)

	select {
	case res := <-prodChan:
		if len(res) != 2 {
			t.Errorf("expected 2 products, got %d", len(res))
		}
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestProductRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewProductRepository(dao, queries.NewProductQueries())

	product := &domain.Product{
		Id: "prod-1", CompanyId: "comp-1", BranchId: "branch-1", BarCode: "123-updated",
		Name: "Prod Updated", Description: "Desc Updated", Discount: 10,
		ProfitPercent: 20, Origin: domain.Imported, UpdatedAt: time.Now(),
	}

	mock.ExpectExec("UPDATE products").
		WithArgs(product.BarCode, product.Name, product.Description, product.Discount,
			product.ProfitPercent, int(product.Origin), product.UpdatedAt, product.CompanyId,
			product.BranchId, product.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errChan := repo.Update(product)

	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("test timed out")
	}
}

func TestProductRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dao := &postgres.DAO{DB: db}
	repo := NewProductRepository(dao, queries.NewProductQueries())

	id := "prod-1"
	compId := "comp-1"
	branchId := "branch-1"
	now := time.Now()

	mock.ExpectExec("UPDATE products").
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
