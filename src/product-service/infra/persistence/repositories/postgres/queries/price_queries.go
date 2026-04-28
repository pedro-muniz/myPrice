package queries

type PriceQueries struct{}

func NewPriceQueries() *PriceQueries {
	return &PriceQueries{}
}

func (q *PriceQueries) Insert() string {
	return "INSERT INTO prices (company_id, branch_id, gross, net, selling, recommended, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
}
func (q *PriceQueries) FindById() string {
	return "SELECT id, company_id, branch_id, gross, net, selling, recommended, created_at, updated_at, deleted_at FROM prices WHERE company_id = $1 AND branch_id = $2 AND id = $3 AND deleted_at IS NULL;"
}
func (q *PriceQueries) FindAll() string {
	return "SELECT id, company_id, branch_id, gross, net, selling, recommended, created_at, updated_at, deleted_at FROM prices WHERE company_id = $1 AND branch_id = $2 AND deleted_at IS NULL;"
}
func (q *PriceQueries) Update() string {
	return "UPDATE prices SET gross = $1, net = $2, selling = $3, recommended = $4, updated_at = $5 WHERE company_id = $6 AND branch_id = $7 AND id = $8;"
}
func (q *PriceQueries) Delete() string {
	return "UPDATE prices SET deleted_at = $1 WHERE company_id = $2 AND branch_id = $3 AND id = $4;"
}
