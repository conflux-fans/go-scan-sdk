package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountTransactions(t *testing.T) {
	client := NewClient("https://api.confluxscan.io")
	start, err := time.ParseInLocation(time.DateOnly, "2024-10-01", time.Local)
	assert.NoError(t, err)
	end, err := time.ParseInLocation(time.DateOnly, "2024-11-01", time.Local)
	assert.NoError(t, err)
	transactions, err := client.GetAccountTransactions("cfx:aakk91pj0pzcbrjkefttdf27t072f4u8pj27znjbw0", "", "", start, end, 0, 100, true)
	assert.Nil(t, err)
	assert.NotNil(t, transactions)
	fmt.Println(transactions.Total)
}

func TestUrl(t *testing.T) {
	client := NewClient("https://api.confluxscan.io")
	assert.Equal(t, "https://api.confluxscan.io/account/transactions", client.url("/account/transactions"))
}
