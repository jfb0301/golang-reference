package main

import(
	"fmt"
)

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.Begin(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	} ()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (VAL) values $1", value1)
	if err != nil {
		return err
	}
	// use tx to do more databases inserts here 
	return nil 
}