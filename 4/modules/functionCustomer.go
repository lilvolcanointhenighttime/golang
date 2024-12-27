package modules

import (
	"context"
	// "database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Customer struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	Birthdate *time.Time `json:"birthDate"`
}

func (db *DB) GetCustomer(email, password string) (Customer, error) {
	ctx := context.Background()

	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return Customer{}, err
	}
	defer conn.Release()

	user := Customer{}

	_ = conn.QueryRow(ctx,
		`SELECT * FROM customers WHERE email = $1`, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.Birthdate,
	)
	fmt.Println(user)
	isEqual := CheckPasswordHash(password, user.Password)
	fmt.Println(isEqual)
	if !isEqual {
		return Customer{}, errors.New("password is incorrect")
	}
	return user, nil
}

func (db *DB) CheckCustomer(user Customer) error {
	ctx := context.Background()

	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx,
		`SELECT id, email, password, name, surname, birthdate FROM customers WHERE email = $1`, user.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.Birthdate,
	)
	if err.Error() != "no rows in result set" {
		fmt.Println(err)
		fmt.Println("48r54i548iti54ughi4u5hut")
		return err
	}
	return nil
}

func (db *DB) CreateCustomer(user Customer) (Customer, error) {
	ctx := context.Background()
	fmt.Println("1")

	conn, err := db.pool.Acquire(ctx)
	fmt.Println("2")
	if err != nil {
		fmt.Println(err)
		return Customer{}, err
	}
	defer conn.Release()

	fmt.Println("3")
	err = conn.QueryRow(ctx,
		`INSERT INTO customers (email, password, name, surname, birthdate) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		user.Email, user.Password, user.Name, user.Surname, user.Birthdate).Scan(&user.ID)
	fmt.Println(err)
	fmt.Println("4")

	return user, err
}

func (db *DB) CreateTableCustomers() error {
	ctx := context.Background()

	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx,
		`CREATE TABLE customers (id SERIAL PRIMARY KEY, email varchar(254), password varchar(254), name TEXT NOT NULL, surname TEXT, birthdate TIMESTAMP)`,
	).Scan()

	return err
}
