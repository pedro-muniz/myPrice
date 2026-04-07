package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	repository "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/auth"
)

type DAO struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

var (
	instance *DAO
	once     sync.Once
)

func GetInstance(host string, port int, user, password, dbName string) *DAO {
	once.Do(func() {
		instance = &DAO{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			DbName:   dbName,
		}
	})
	return instance
}

func (this *DAO) getConnection() (*sql.DB, error) {
	fmt.Println("Connecting to postgres")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		this.Host, this.Port, this.User, this.Password, this.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging postgres: %v", err)
	}

	fmt.Println("connected to postgres")
	return db, nil
}

func (this *DAO) Insert(authModel repository.AuthModel) (sql.Result, error) {
	db, err := this.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("writing to postgres")

	sqlStatement := `
	INSERT INTO user_logins.auths (name, email, password, roles, created_at, 
	updated_at, last_login_at)
	VALUES ($1, $2, $3, $4, NOW(), NOW(), NOW())`

	result, err := db.Exec(sqlStatement, authModel.Name, authModel.Email, authModel.Password, authModel.Roles)
	if err != nil {
		return nil, fmt.Errorf("error running create: %v", err)
	}

	return result, nil
}

func (this *DAO) Read(query string, args ...interface{}) (*sql.Rows, error) {
	db, err := this.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("reading from postgres")

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error running read query: %v", err)
	}

	return rows, nil
}

func (this *DAO) Update(authModel repository.AuthModel) (sql.Result, error) {
	db, err := this.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("updating in postgres")

	sqlStatement := `
	UPDATE user_logins.auths 
	SET name = $2, email = $3, password = $4, roles = $5, updated_at = NOW(), last_login_at = $6
	WHERE auth_id = $1`

	result, err := db.Exec(sqlStatement, authModel.Id, authModel.Name, authModel.Email, authModel.Password, authModel.Roles, authModel.LastLoginAt)
	if err != nil {
		return nil, fmt.Errorf("error running update: %v", err)
	}

	return result, nil
}

func (this *DAO) Delete(id string) (sql.Result, error) {
	db, err := this.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("deleting from postgres")

	sqlStatement := `DELETE FROM user_logins.auths WHERE auth_id = $1`

	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		return nil, fmt.Errorf("error running delete: %v", err)
	}

	return result, nil
}
