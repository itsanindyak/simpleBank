package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/itsanindyak/simpleBank/utils"
	"github.com/stretchr/testify/require"
)
func createRandomAccount(t *testing.T) Accounts{
	arg:= CreateAccountParams{
		Owner: utils.RandomOwner(),
		Balance:utils.RandomMoney(),
		Currency: "INR",
	}

	account, err := testQuires.CreateAccount(context.Background(),arg)

	require.NoError(t,err)
	require.NotEmpty(t,account)

	require.Equal(t,account.Owner,arg.Owner)
	require.Equal(t,account.Balance,arg.Balance)
	require.Equal(t,account.Currency,arg.Currency)

	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)

	return *account


}
func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T){
	account1 := createRandomAccount(t)

	account2,err := testQuires.GetAccount(context.Background(),account1.ID)

	// test
	require.NoError(t,err)
	require.NotEmpty(t,account2)
	
	require.Equal(t,account1.ID,account2.ID)
	require.Equal(t,account1.Owner,account2.Owner)
	require.Equal(t,account1.Balance,account2.Balance)
	require.Equal(t,account1.Currency,account2.Currency)
}


func TestUpdateAccount(t *testing.T){
	account := createRandomAccount(t)
	arg:= UpdateAccountParams{
		ID: account.ID ,
		Balance: utils.RandomMoney(),
	}

	updatedAccount,err := testQuires.UpdateAccount(context.Background(),arg)

	// test
	require.NoError(t,err)
	require.NotEmpty(t,updatedAccount)


	require.Equal(t,account.ID,updatedAccount.ID)
	require.Equal(t,account.Owner,updatedAccount.Owner)
	require.Equal(t,arg.Balance,updatedAccount.Balance)
	require.Equal(t,account.Currency,updatedAccount.Currency)
}

func TestDeleteAccount(t *testing.T){
	account1 := createRandomAccount(t)

	err := testQuires.DeleteAccount(context.Background(),account1.ID)

	require.NoError(t,err)

	account2,err := testQuires.GetAccount(context.Background(),account1.ID)
	require.Error(t,err)
	require.EqualError(t,err,sql.ErrNoRows.Error())
	require.Empty(t,account2)

}

func TestListAccount(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts ,err :=testQuires.ListAccounts(context.Background(),arg)


	require.NoError(t,err)
	require.Len(t,accounts,5)

	for _,account := range accounts{
		require.NotEmpty(t,account)
	}
}