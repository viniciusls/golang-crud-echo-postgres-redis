package dao

import (
	"crud-echo-postgres-redis/config"
	"crud-echo-postgres-redis/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func createConnection() *sql.DB {
	env, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading app.env file")
	}

	db, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

func GetAllUsers() ([]models.User, error) {
	db := createConnection()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Unable to close database connection")
		}
	}(db)

	var users []models.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Unable to close rows pointer")
		}
	}(rows)

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Location)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)
	}

	return users, err
}

func GetUser(id int64) (models.User, error) {
	db := createConnection()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Unable to close database connection")
		}
	}(db)

	var user models.User

	sqlStatement := `SELECT * FROM users WHERE id = $1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.Id, &user.Name, &user.Age, &user.Location)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")

		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

func CreateUser(user *models.User) int64 {
	db := createConnection()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Unable to close database connection")
		}
	}(db)

	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v\n", id)

	return id
}

func UpdateUser(id int64, user *models.User) int64 {
	db := createConnection()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Unable to close database connection")
		}
	}(db)

	sqlStatement := `UPDATE users SET name = $2, location = $3, age = $4 WHERE id = $1`

	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/records affected %v\n", rowsAffected)

	return rowsAffected
}

func DeleteUser(id int64) int64 {
	db := createConnection()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Unable to close database connection")
		}
	}(db)

	sqlStatement := `DELETE FROM users WHERE id = $1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/records affected %v\n", rowsAffected)

	return rowsAffected
}
