package queries

type CostQueries struct{}

func NewCostQueries() *CostQueries {
	return &CostQueries{}
}

func (q *CostQueries) Insert() string {
	return "INSERT INTO costs (company_id, branch_id, name, value) VALUES ($1, $2, $3, $4);"
}
func (q *CostQueries) FindById() string {
	return "SELECT id, company_id, branch_id, name, value, created_at, updated_at, deleted_at FROM costs WHERE company_id = $1 AND branch_id = $2 AND id = $3 AND deleted_at IS NULL;"
}
func (q *CostQueries) FindAll() string {
	return "SELECT id, company_id, branch_id, name, value, created_at, updated_at, deleted_at FROM costs WHERE company_id = $1 AND branch_id = $2 AND deleted_at IS NULL;"
}
func (q *CostQueries) Update() string {
	return "UPDATE costs SET name = $1, value = $2, updated_at = $3 WHERE company_id = $4 AND branch_id = $5 AND id = $6;"
}
func (q *CostQueries) Delete() string {
	return "UPDATE costs SET deleted_at = $1 WHERE company_id = $2 AND branch_id = $3 AND id = $4;"
}
