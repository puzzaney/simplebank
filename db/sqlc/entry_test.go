package db

import (
	"context"
	"testing"
	"time"

	"github.com/puzzaney/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(1, 100),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, args.AccountID)
	require.Equal(t, entry.Amount, args.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)

	CreateRandomEntry(t, account)

}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := CreateRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1, entry2)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, arg.AccountID)
	}
}
