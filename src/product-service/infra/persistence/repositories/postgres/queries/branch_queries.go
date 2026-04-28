package queries

type BranchQueries struct{}

func NewBranchQueries() *BranchQueries {
	return &BranchQueries{}
}

func (q *BranchQueries) Insert() string {
	return "INSERT INTO branches (company_id, name) VALUES ($1, $2);"
}
func (q *BranchQueries) FindById() string {
	return "SELECT id, company_id, name, deleted_at FROM branches WHERE company_id = $1 AND id = $2 AND deleted_at IS NULL;"
}
func (q *BranchQueries) FindAll() string {
	return "SELECT id, company_id, name, deleted_at FROM branches WHERE company_id = $1 AND deleted_at IS NULL;"
}
func (q *BranchQueries) Update() string {
	return "UPDATE branches SET name = $1 WHERE company_id = $2 AND id = $3;"
}
func (q *BranchQueries) Delete() string {
	return "UPDATE branches SET deleted_at = $1 WHERE company_id = $2 AND id = $3;"
}
