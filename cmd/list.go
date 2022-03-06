package main

import (
	"context"
	"fmt"

	"github.com/arganaphangquestian/eaas/data"
	"github.com/davecgh/go-spew/spew"
)

func list(ctx context.Context, queries *data.Queries) {
	rows, err := queries.GetCustomers(ctx)
	if err != nil {
		fmt.Println("Failed to get Customers data")
		return
	}
	for row := range rows {
		spew.Print(row)
	}
}
