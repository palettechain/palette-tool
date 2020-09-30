package main

import (
	"math/big"
	"time"

	"github.com/ipfs/go-log"
	"github.com/palettechain/palette-tool/config"
	"github.com/palettechain/palette-tool/utils"
)

const cfgPath = "config.toml"

var (
	logger = log.Logger("geth")
	cfg    = config.GenerateConfig(cfgPath)
	c      = utils.NewPaletteClient(cfg)
)

func main() {
	// testMint()
	testTransfer()
	// testApprove()
	// testTotalSupply(c)
	// testDecimals()
}

//func testMint() {
//	var amount int64 = 10000
//
//	data := new(big.Int).Mul(big.NewInt(amount), utils.OnePLT)
//	hash, err := c.PLTMint(data)
//	if err != nil {
//		logger.Fatal(err)
//	}
//	logger.Infof("mint hash %s", hash.Hex())
//
//	waitCommit()
//
//	if err := c.DumpEventLog(hash); err != nil {
//		logger.Fatal(err)
//	}
//}

// 管理员给测试用户转账10PLT
func testTransfer() {
	var num int64 = 10

	amount := new(big.Int).Mul(big.NewInt(num), utils.OnePLT)
	to := c.TestAccounts[0].Address

	hash, err := c.PLTTransfer(c.Admin.PrivateKey, to, amount)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("transfer hash %s", hash.Hex())

	waitCommit()

	if err := c.DumpEventLog(hash); err != nil {
		logger.Fatal(err)
	}
}

// 测试账户0给测试账户1授权10PLT
func testApprove() {
	var num int64 = 10

	owner := c.TestAccounts[0].PrivateKey
	spender := c.TestAccounts[1].Address
	amount := new(big.Int).Mul(big.NewInt(num), utils.OnePLT)

	hash, err := c.PLTApprove(owner, spender, amount)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("approve hash %s", hash.Hex())

	waitCommit()

	if err := c.DumpEventLog(hash); err != nil {
		logger.Fatal(err)
	}
}

func testTotalSupply() {
	amount, err := c.PLTTotalSupply()
	if err != nil {
		logger.Fatal(err)
	}

	supply := new(big.Int).Div(amount, utils.OnePLT)
	logger.Infof("PLT total supply is %s", supply.String())
}

func testDecimals() {
	decimals, err := c.PLTDecimals()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("PLT decimals %d", decimals)
}

func waitCommit() {
	time.Sleep(18 * time.Second)
}
