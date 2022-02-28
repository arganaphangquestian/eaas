package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/arganaphangquestian/eaas/data"
	"github.com/davecgh/go-spew/spew"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type Customer struct {
	ID      string
	Name    string
	Email   string
	Address string
	Balance string
}

var (
	DATABASE_URL                   string
	name, email, password, address string
	balance                        int
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Can't Load .env Config")
	}
	DATABASE_URL = fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
}

func add(ctx context.Context, queries *data.Queries) {
	flag.StringVar(&name, "name", "", "Customer name")
	flag.StringVar(&email, "email", "", "Customer email")
	flag.StringVar(&password, "password", "", "Customer password")
	flag.StringVar(&address, "address", "", "Customer address")
	flag.IntVar(&balance, "balance", 0, "Customer balance")
	flag.Parse()
	balance_numeric := pgtype.Numeric{}
	balance_numeric.Set(balance)
	customer, err := queries.AddCustomer(ctx, data.AddCustomerParams{
		Name:     name,
		Email:    email,
		Password: sql.NullString{String: password, Valid: true},
		Balance:  balance_numeric,
	})
	if err != nil {
		log.Println("Failed to add new Customer!")
		return
	}

	log.Println("New Customer added!")
	spew.Dump(Customer{
		ID:      customer.ID.String(),
		Name:    customer.Name,
		Email:   customer.Email,
		Address: customer.Address.String,
		Balance: customer.Balance.Int.String(),
	})
}

func list(ctx context.Context, queries *data.Queries) {
	customers, err := queries.GetCustomers(ctx)
	if err != nil {
		log.Println("Failed to retrive Customer Table")
	}
	for _, v := range customers {
		spew.Dump(Customer{
			ID:      v.ID.String(),
			Name:    v.Name,
			Email:   v.Email,
			Address: v.Address.String,
			Balance: v.Balance.Int.String(),
		})
	}
}

func main() {
	// Database Connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	defer conn.Close(ctx)
	queries := data.New(conn)

	if command == "add" {
		add(ctx, queries)
		return
	}

	list(ctx, queries)
}
