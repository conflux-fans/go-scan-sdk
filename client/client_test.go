package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

func TestGetPosAccountOverview(t *testing.T) {
	client := NewClient("https://confluxscan.io")
	overview, err := client.GetPosAccountOverview(common.HexToHash("0x6aab785e2f7bc3656825ae1b674e7ec9159e573326e5b0f5acf4f1ed46ace34d"))
	assert.Nil(t, err)
	assert.NotNil(t, overview)
	fmt.Println(overview)
}

func TestGetPosAccountReward(t *testing.T) {
	client := NewClient("https://confluxscan.io")
	reward, err := client.GetPosAccountReward(common.HexToHash("0xae888cc930f28bd81c22f3783f615d03701363a06ad24b90aca5ef5a15d758b0"), "incoming-history")
	assert.Nil(t, err)
	assert.NotNil(t, reward)
	fmt.Println(reward.Total)
	for _, r := range reward.List {
		fmt.Println(r.ID, r.AccountID, r.Reward, r.CreatedAt, r.Epoch, r.PowBlockHash.Hex())
	}
}

func TestUrl(t *testing.T) {
	client := NewClient("https://api.confluxscan.io")
	assert.Equal(t, "https://api.confluxscan.io/account/transactions", client.url("/account/transactions"))
}
