package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arganaphangquestian/eaas/data"
	"github.com/jackc/pgtype"
)

func seed(ctx context.Context, queries *data.Queries, rows int) error {
	for i := 0; i < rows; i++ {
		row := i + 1
		_, err := queries.AddCustomer(ctx, data.AddCustomerParams{
			Name:  fmt.Sprintf("Name %d", row),
			Email: fmt.Sprintf("email%d@mail.com", row),
			Password: sql.NullString{
				String: fmt.Sprintf("password%d", row),
				Valid:  true,
			},
			Address: sql.NullString{
				String: fmt.Sprintf("address ke %d", row),
				Valid:  true,
			},
			Balance: pgtype.Numeric{
				Exp: int32(row) * 1000,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
