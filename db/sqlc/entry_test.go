package db

import (
	"context"
	"testing"

	"github.com/itsanindyak/simpleBank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T,account Accounts) Entries {
	arg:= CreateEntryParams{
		AccountsID: account.ID,
		Amount: utils.RandomMoney(),
	}

	entry,err := testQuires.CreateEntry(context.Background(),arg)

	require.NoError(t,err)
	require.NotEmpty(t,entry)

	require.Equal(t,account.ID,entry.AccountsID)
	require.Equal(t,entry.Amount,arg.Amount)

	require.NotZero(t,entry.ID)
	require.NotZero(t,entry.CreatedAt)

	return *entry

}

func TestCreateEntry(t *testing.T){
	account := createRandomAccount(t)
	createRandomEntry(t,account)
}

func TestGetEntry(t *testing.T){
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t,account)

	entry2,err := testQuires.GetEntry(context.Background(),entry1.ID)

	require.NoError(t,err)
	require.NotEmpty(t,entry2)

	require.Equal(t,entry1.ID,entry2.ID)
	require.Equal(t,entry1.Amount,entry2.Amount)
	require.Equal(t,entry1.AccountsID,entry2.AccountsID)
}


func TestListEntry(t *testing.T){
	account := createRandomAccount(t)
	for range 10{
		createRandomEntry(t,account)
	}

	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries ,err :=testQuires.ListEntries(context.Background(),arg)


	require.NoError(t,err)
	require.Len(t,entries,5)

	for _,entry := range entries{
		require.NotEmpty(t,entry)
	}
}