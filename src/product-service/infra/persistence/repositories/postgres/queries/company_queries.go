package queries

type CompanyQueries struct{}

func NewCompanyQueries() *CompanyQueries {
	return &CompanyQueries{}
}

func (q *CompanyQueries) Insert() string { return "INSERT INTO companies (name) VALUES ($1);" }
func (q *CompanyQueries) FindById() string {
	return "SELECT id, name, deleted_at FROM companies WHERE id = $1 AND deleted_at IS NULL;"
}
func (q *CompanyQueries) FindAll() string {
	return "SELECT id, name, deleted_at FROM companies WHERE deleted_at IS NULL;"
}
func (q *CompanyQueries) Update() string { return "UPDATE companies SET name = $1 WHERE id = $2;" }
func (q *CompanyQueries) Delete() string {
	return "UPDATE companies SET deleted_at = $1 WHERE id = $2;"
}
