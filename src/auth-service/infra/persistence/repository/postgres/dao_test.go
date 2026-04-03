package postgres

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	repository "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/auth"
)

// Warn: This test relies on database connection and runs sequentially
func TestDAOLifecycle_validData_shouldPass(t *testing.T) {
	port := 5432
	if envPort := os.Getenv("postgres_port"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		}
	}

	var dao *DAO = &DAO{
		Host:     os.Getenv("postgres_host"),
		Port:     port,
		User:     os.Getenv("postgres_user"),
		Password: os.Getenv("postgres_pass"),
		DbName:   os.Getenv("postgres_db"),
	}

	if dao.Host == "" {
		t.Skip("Skipping postgres DAO tests because postgres_host env var is not set")
	}

	var generatedId string

	t.Run("Insert", func(t *testing.T) {
		authModel := repository.AuthModel{
			Name:     "Mocked Postgres User",
			Email:    "mocked_pg@email.com",
			Password: "testing password",
			Roles:    "admin",
		}

		result, err := dao.Insert(authModel)
		if err != nil {
			t.Fatalf("The insert method didn't return the expected result: %v", err)
		}
		fmt.Println("Insert Result:", result)

		// Retrieve the generated ID
		rows, err := dao.Read("SELECT auth_id FROM user_logins.auths WHERE email = $1 ORDER BY created_at DESC LIMIT 1", authModel.Email)
		if err != nil {
			t.Fatalf("Failed to read back the inserted id: %v", err)
		}
		defer rows.Close()

		if rows.Next() {
			rows.Scan(&generatedId)
		}

		if generatedId == "" {
			t.Fatalf("Generated Id was empty after insertion")
		}
		fmt.Println("Generated DB ID:", generatedId)
	})

	t.Run("Read", func(t *testing.T) {
		rows, err := dao.Read("SELECT auth_id, name, email FROM user_logins.auths WHERE auth_id = $1", generatedId)
		if err != nil {
			t.Fatalf("The read method didn't return the expected result: %v", err)
		}
		defer rows.Close()
		fmt.Println("Successfully read inserted ID from Postgres DB")
	})

	t.Run("Update", func(t *testing.T) {
		authModel := repository.AuthModel{
			Id:       generatedId,
			Name:     "Updated Mocked User",
			Email:    "updated_mocked@email.com",
			Password: "new password",
			Roles:    "user",
		}

		result, err := dao.Update(authModel)
		if err != nil {
			t.Fatalf("The update method didn't return the expected result: %v", err)
		}
		fmt.Println("Update Result:", result)
	})

	t.Run("Delete", func(t *testing.T) {
		result, err := dao.Delete(generatedId)
		if err != nil {
			t.Fatalf("The delete method didn't return the expected result: %v", err)
		}
		fmt.Println("Delete Result:", result)
	})
}
