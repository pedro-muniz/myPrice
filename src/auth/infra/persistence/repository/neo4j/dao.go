//package to handle with neo4j calls - https://pkg.go.dev/github.com/neo4j/neo4j-go-driver/v4#readme-documentation
package neo4j

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	repository "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/neo4j/auth"
)

type DAO struct {
	Neo4jUser     string
	Neo4JPassword string
}

func (this *DAO) getDriver() (neo4j.Driver, error) {
	fmt.Println("Connecting to neo4j")
	driver, err := neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth(this.Neo4jUser, this.Neo4JPassword, ""))
	if err != nil {
		fmt.Errorf("Error connecting to neo4j: %v", err)
		return nil, err
	}
	fmt.Println("connected to neo4j")

	return driver, err
}

func (this *DAO) Write(authModel repository.AuthModel) (neo4j.Result, error) {
	driver, err := this.getDriver()
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	fmt.Println("writing to neo4j")
	records, err := session.Run("CREATE (a:Auth {name: $name, email: $email, password: $password, roles: $roles})",
		map[string]interface{}{
			"name":     authModel.Name,
			"email":    authModel.Email,
			"password": authModel.Password,
			"roles":    authModel.Roles,
		})

	if err != nil {
		fmt.Errorf("Error running create: %v", err)
		return nil, err
	}

	return records, nil
}
