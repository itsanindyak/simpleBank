package db

import (
	"context"
	"testing"

	"github.com/itsanindyak/simpleBank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, from, to Accounts) Transfers {
	arg := CreateTransferParams{
		FromAccountsID: from.ID,
		ToAccountsID:   to.ID,
		Amount:         utils.RandomMoney(),
	}

	transfer, err := testQuires.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountsID, transfer.FromAccountsID)
	require.Equal(t, arg.ToAccountsID, transfer.ToAccountsID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return *transfer

}

func TestCreateTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	createRandomTransfer(t, from, to)

}


func TestGetTransfer(t *testing.T){
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	transfers1 := createRandomTransfer(t, from, to)

	transfers2,err := testQuires.GetTransfer(context.Background(),transfers1.ID)

	require.NoError(t,err)
	require.NotEmpty(t,transfers2)

	require.Equal(t,transfers1.ID,transfers2.ID)
	require.Equal(t,transfers1.FromAccountsID,transfers2.FromAccountsID)
	require.Equal(t,transfers1.ToAccountsID,transfers2.ToAccountsID)
	require.Equal(t,transfers1.Amount,transfers2.Amount)
}

func TestListTransfer(t *testing.T){
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	for range 10{
		createRandomTransfer(t,from,to)
	}

	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	transfers ,err :=testQuires.ListTransfers(context.Background(),arg)


	require.NoError(t,err)
	require.Len(t,transfers,5)

	for _,transfer := range transfers{
		require.NotEmpty(t,transfer)
	}
}
