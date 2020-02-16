package db

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	connStr             = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	createEmployeeTable = `CREATE TABLE IF NOT EXISTS employee (id uuid PRIMARY KEY,
                                             name varchar,
                                             age int,
                                             language varchar)`

	insertEmployee = "INSERT INTO employee(id, name, age, language) VALUES ($1, $2, $3, $4)"
	selectEmployee = `SELECT name, age, language FROM employee WHERE id = $1`
)

type (
	YB struct {
		db *sql.DB
	}
)

func NewYB(port int, host, user, password, dbname string) (*YB, error) {
	var (
		err error

		yb       = YB{}
		psqlInfo = fmt.Sprintf(connStr, host, port, user, password, dbname)
	)

	yb.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	_, err = yb.db.Exec(createEmployeeTable)
	if err != nil {
		return nil, err
	}

	fmt.Println("Created table employee")

	return &yb, nil
}

func (yb *YB) NewEmployee(e *Employee) (uuid.UUID, error) {
	var id = uuid.New()

	// Insert into the table.
	var (
		_, err = yb.db.Exec(insertEmployee, id.String(), e.Name, e.Age, e.Language)
	)

	if err != nil {
		return uuid.Nil, nil
	}

	return id, nil
}

func (yb *YB) GetEmployeeByID(id uuid.UUID) (*Employee, error) {
	// Read from the table.
	var (
		e = Employee{
			ID: id,
		}

		rows, err = yb.db.Query(selectEmployee, id.String())
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&e.Name, &e.Age, &e.Language)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (yb *YB) Close() error {
	return yb.db.Close()
}
