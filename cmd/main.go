package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/arganaphangquestian/eaas/data"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var (
	DATABASE_URL string
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Can't Load .env Config")
	}
	DATABASE_URL = fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
}

func main() {
	// Args
	var name, email, password, address string
	var balance int
	flag.StringVar(&name, "name", "", "Customer name")
	flag.StringVar(&email, "email", "", "Customer email")
	flag.StringVar(&password, "password", "", "Customer password")
	flag.StringVar(&address, "address", "", "Customer address")
	flag.IntVar(&balance, "balance", 0, "Customer balance")
	flag.Parse()

	// Processing
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	defer conn.Close(ctx)
	queries := data.New(conn)

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
	log.Printf("%#v\n", customer)
}
