package main

import (
	"context"
	"fmt"

	"github.com/arganaphangquestian/eaas/data"
)

func add(ctx context.Context, queries *data.Queries, input data.AddCustomerParams) {
	_, err := queries.AddCustomer(ctx, input)
	if err != nil {
		fmt.Println("Insert Customer Data Failed!")
		return
	}
	fmt.Println("Insert Customer Data Success!")
}
