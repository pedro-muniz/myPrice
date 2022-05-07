package neo4j

import (
	"fmt"
	"os"
	"testing"

	repository "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/neo4j/auth"
)

//Warn: This test is not mocking neo4j driver and session
func TestDAOWrite_validData_shouldPass(t *testing.T) {
	//arrange
	var authModel repository.AuthModel = repository.AuthModel{
		Name:     "Pedro",
		Email:    "pmuniz09@gmail.com",
		Password: "testing password",
		Roles:    "admin",
	}

	var dao *DAO = &DAO{
		Neo4jUser:     os.Getenv("neo4jUser"),
		Neo4JPassword: os.Getenv("neo4jPass"),
	}

	records, err := dao.Write(authModel)
	fmt.Println(records.Record())
	if err != nil {
		fmt.Println(err)
		t.Error("The authorize didn't return the expectedError")
	}
}
