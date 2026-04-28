package queries

type ProductQueries struct{}

func NewProductQueries() *ProductQueries {
	return &ProductQueries{}
}

func (q *ProductQueries) Insert() string {
	return "INSERT INTO products (company_id, branch_id, barcode, name, description, discount, price_id, profit_percent, origin, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
}
func (q *ProductQueries) FindById() string {
	return "SELECT pd.id, pd.company_id, pd.branch_id, pd.barcode, pd.name, pd.description, pd.discount, pd.profit_percent, pd.origin, pd.created_at, pd.updated_at, pd.deleted_at, pr.id, pr.gross, pr.net, pr.selling, pr.recommended, pr.created_at, pr.updated_at, pr.deleted_at FROM products pd JOIN prices pr ON pd.price_id = pr.id WHERE pd.company_id = $1 AND pd.branch_id = $2 AND pd.id = $3 AND pd.deleted_at IS NULL;"
}
func (q *ProductQueries) FindAll() string {
	return "SELECT pd.id, pd.company_id, pd.branch_id, pd.barcode, pd.name, pd.description, pd.discount, pd.profit_percent, pd.origin, pd.created_at, pd.updated_at, pd.deleted_at, pr.id, pr.gross, pr.net, pr.selling, pr.recommended, pr.created_at, pr.updated_at, pr.deleted_at FROM products pd JOIN prices pr ON pd.price_id = pr.id WHERE pd.company_id = $1 AND pd.branch_id = $2 AND pd.deleted_at IS NULL;"
}
func (q *ProductQueries) Update() string {
	return "UPDATE products SET barcode = $1, name = $2, description = $3, discount = $4, profit_percent = $5, origin = $6, updated_at = $7 WHERE company_id = $8 AND branch_id = $9 AND id = $10;"
}
func (q *ProductQueries) Delete() string {
	return "UPDATE products SET deleted_at = $1 WHERE company_id = $2 AND branch_id = $3 AND id = $4;"
}
