package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v and rb err %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()

}

// Transfer money one account to another account
// transfer record --> entry record in to_account ---> entry record in from_account ---> update account balance in both account

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    *Transfers `json:"transfer"`
	FromAccount *Accounts  `json:"from_account"`
	ToAccount   *Accounts  `json:"to_account"`
	FromEntry   *Entries   `json:"from_entry"`
	ToEntry     *Entries   `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// create transfer

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountsID: arg.FromAccountId,
			ToAccountsID:   arg.ToAccountId,
			Amount:         arg.Amount,
		})

		if err != nil {
			return err
		}

		// create entry with -ve for from account

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountsID: arg.FromAccountId,
			Amount:     -arg.Amount,
		})

		if err != nil {
			return err
		}
		// create entry with +ve for to_account

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountsID: arg.ToAccountId,
			Amount:     arg.Amount,
		})

		if err != nil {
			return err
		}

		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount ,result.ToAccount,err =  addMoney(ctx,q,arg.FromAccountId,-arg.Amount,arg.ToAccountId,arg.Amount)

		}else{
			result.ToAccount ,result.FromAccount,err =  addMoney(ctx,q,arg.ToAccountId,arg.Amount,arg.FromAccountId,-arg.Amount)


		}
		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountId1 int64,
	amount1 int64,
	accountId2 int64,
	amount2 int64,
)(account1 *Accounts, account2 *Accounts, err error) {
	account1, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		Amount: amount1,
		ID:     accountId1,
	})
	if err != nil{
		return
	}
	account2, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		Amount: amount2,
		ID:     accountId2,
	})
	if err != nil{
		return
	}
	return
}