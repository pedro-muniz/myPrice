package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type DAO struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	DB       *sql.DB
}

var (
	instance *DAO
	once     sync.Once
)

func GetInstance(host string, port int, user, password, dbName string) *DAO {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbName)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Printf("Error opening database connection pool: %v\n", err)
		} else {
			// Test the connection
			if err := db.Ping(); err != nil {
				fmt.Printf("Error pinging database: %v\n", err)
			}
		}

		instance = &DAO{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			DbName:   dbName,
			DB:       db,
		}
	})
	return instance
}

// Write supports INSERT, UPDATE, and DELETE queries
func (this *DAO) Write(query string, args ...interface{}) (sql.Result, error) {
	fmt.Println("writing to postgres")

	if this.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	result, err := this.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error running write query: %v", err)
	}

	return result, nil
}

// Read supports SELECT queries
func (this *DAO) Read(query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println("reading from postgres")

	if this.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	rows, err := this.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error running read query: %v", err)
	}

	return rows, nil
}
