package client

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type List[T any] struct {
	Total int `json:"total"`
	List  []T `json:"list"`
}

type Transaction struct {
	EpochNumber      int     `json:"epochNumber"`
	BlockPosition    int     `json:"blockPosition"`
	TransactionIndex int     `json:"transactionIndex"`
	Nonce            string  `json:"nonce"`
	Hash             string  `json:"hash"`
	From             string  `json:"from"`
	To               string  `json:"to"`
	Value            string  `json:"value"`
	GasPrice         string  `json:"gasPrice"`
	GasFee           string  `json:"gasFee"`
	Timestamp        int     `json:"timestamp"`
	Status           int     `json:"status"`
	ContractCreated  *string `json:"contractCreated"`
	Method           string  `json:"method"`
}

type PosAccountOverview struct {
	Address             common.Hash `json:"address"`
	Byte32NameTagInfo   *string     `json:"byte32NameTagInfo"`
	CreatedAt           time.Time   `json:"createdAt"`
	Type                string      `json:"type"`
	Status              string      `json:"status"`
	AvailableVotesInCfx int         `json:"availableVotesInCfx"`
	WithdrawableInCfx   int         `json:"withdrawableInCfx"`
	LockingInCfx        int         `json:"lockingInCfx"`
	UnlockingInCfx      int         `json:"unlockingInCfx"`
	ForceRetired        *string     `json:"forceRetired"`
	TotalReward         string      `json:"totalReward"`
}
