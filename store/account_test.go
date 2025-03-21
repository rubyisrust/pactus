package store

import (
	"testing"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/account"
	"github.com/pactus-project/pactus/util/testsuite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountCounter(t *testing.T) {
	td := setup(t, nil)

	acc, addr := td.GenerateTestAccount()

	t.Run("Add new account, should increase the total accounts number", func(t *testing.T) {
		assert.Zero(t, td.store.TotalAccounts())

		td.store.UpdateAccount(addr, acc)
		assert.NoError(t, td.store.WriteBatch())
		assert.Equal(t, int32(1), td.store.TotalAccounts())
	})

	t.Run("Update account, should not increase the total accounts number", func(t *testing.T) {
		acc.AddToBalance(1)
		td.store.UpdateAccount(addr, acc)

		assert.NoError(t, td.store.WriteBatch())
		assert.Equal(t, int32(1), td.store.TotalAccounts())
	})

	t.Run("Get account", func(t *testing.T) {
		acc1, err := td.store.Account(addr)
		assert.NoError(t, err)

		assert.Equal(t, acc1.Hash(), acc.Hash())
		assert.Equal(t, int32(1), td.store.TotalAccounts())
		assert.True(t, td.store.HasAccount(addr))
	})
}

func TestAccountBatchSaving(t *testing.T) {
	td := setup(t, nil)

	total := td.RandInt32NonZero(100)
	t.Run("Add some accounts", func(t *testing.T) {
		for i := int32(0); i < total; i++ {
			acc, addr := td.GenerateTestAccount(testsuite.AccountWithNumber(i))
			td.store.UpdateAccount(addr, acc)
		}
		assert.NoError(t, td.store.WriteBatch())
		assert.Equal(t, total, td.store.TotalAccounts())
	})

	t.Run("Close and load db", func(t *testing.T) {
		td.store.Close()
		store, _ := NewStore(td.store.config)
		assert.Equal(t, total, store.TotalAccounts())
	})
}

func TestAccountByAddress(t *testing.T) {
	td := setup(t, nil)

	total := td.RandInt32NonZero(100)
	var lastAddr crypto.Address
	t.Run("Add some accounts", func(t *testing.T) {
		for i := int32(0); i < total; i++ {
			acc, addr := td.GenerateTestAccount(testsuite.AccountWithNumber(i))
			td.store.UpdateAccount(addr, acc)

			lastAddr = addr
		}
		assert.NoError(t, td.store.WriteBatch())
		assert.Equal(t, total, td.store.TotalAccounts())
	})

	t.Run("Non existing account", func(t *testing.T) {
		addr := td.RandAccAddress()
		acc, err := td.store.Account(addr)
		has := td.store.HasAccount(addr)

		assert.False(t, has)
		assert.Error(t, err)
		assert.Nil(t, acc)
	})

	t.Run("Reopen the store", func(t *testing.T) {
		td.store.Close()
		store, _ := NewStore(td.store.config)

		acc, err := store.Account(lastAddr)
		assert.NoError(t, err)
		require.NotNil(t, acc)
		assert.Equal(t, total-1, acc.Number())
	})
}

func TestIterateAccounts(t *testing.T) {
	td := setup(t, nil)

	total := td.RandInt32NonZero(100)
	hashes1 := []hash.Hash{}
	for i := int32(0); i < total; i++ {
		acc, addr := td.GenerateTestAccount(testsuite.AccountWithNumber(i))
		td.store.UpdateAccount(addr, acc)
		hashes1 = append(hashes1, acc.Hash())
	}
	assert.NoError(t, td.store.WriteBatch())

	hashes2 := []hash.Hash{}
	td.store.IterateAccounts(func(_ crypto.Address, acc *account.Account) bool {
		hashes2 = append(hashes2, acc.Hash())

		return false
	})
	assert.ElementsMatch(t, hashes1, hashes2)

	stopped := false
	td.store.IterateAccounts(func(_ crypto.Address, acc *account.Account) bool {
		if acc.Hash() == hashes1[0] {
			stopped = true
		}

		return stopped
	})
	assert.True(t, stopped)
}

func TestAccountDeepCopy(t *testing.T) {
	td := setup(t, nil)

	acc1, addr := td.GenerateTestAccount()
	td.store.UpdateAccount(addr, acc1)

	acc2, _ := td.store.Account(addr)
	acc2.AddToBalance(1)
	accCache, _ := td.store.accountStore.accCache.Get(addr)
	assert.NotEqual(t, acc2.Hash(), accCache.Hash())
}
