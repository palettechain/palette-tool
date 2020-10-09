package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"

	"github.com/ipfs/go-log"
	"github.com/palettechain/palette-tool/config"
	"github.com/palettechain/palette-tool/utils"
)

var (
	logger  = log.Logger("geth")
	cfg    *config.Config
	client *utils.PaletteClient

	cfgpath string
)

func init() {
	flag.StringVar(&cfgpath, "config", "local.toml", "set config path")
}

func main() {
	flag.Parse()
	cfg = config.GenerateConfig(cfgpath)
	client = utils.NewPaletteClient(cfg)

	// testMint()
	testTransfer()
	// testApprove()
	// testTotalSupply(client)
	// testDecimals()
	// testAddValidator()
}

// 管理员给测试用户转账10PLT
func testTransfer() {
	var num int64 = 10

	amount := new(big.Int).Mul(big.NewInt(num), utils.OnePLT)
	to := client.TestAccounts[0].Address

	hash, err := client.PLTTransfer(client.Admin.PrivateKey, to, amount)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("transfer hash %s", hash.Hex())

	waitCommit()

	if err := client.DumpEventLog(hash); err != nil {
		logger.Fatal(err)
	}
}

// 测试账户0给测试账户1授权10PLT
func testApprove() {
	var num int64 = 10

	owner := client.TestAccounts[0].PrivateKey
	spender := client.TestAccounts[1].Address
	amount := new(big.Int).Mul(big.NewInt(num), utils.OnePLT)

	hash, err := client.PLTApprove(owner, spender, amount)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("approve hash %s", hash.Hex())

	waitCommit()

	if err := client.DumpEventLog(hash); err != nil {
		logger.Fatal(err)
	}
}

func testAddValidator() {
	validator := common.HexToAddress("")
	hash, err := client.AddValidator(validator)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("add validator hash %s", hash.Hex())

	waitCommit()

	if err := client.DumpEventLog(hash); err != nil {
		logger.Fatal(err)
	}
}

func testTotalSupply() {
	amount, err := client.PLTTotalSupply()
	if err != nil {
		logger.Fatal(err)
	}

	supply := new(big.Int).Div(amount, utils.OnePLT)
	logger.Infof("PLT total supply is %s", supply.String())
}

func testDecimals() {
	decimals, err := client.PLTDecimals()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("PLT decimals %d", decimals)
}

func waitCommit() {
	time.Sleep(18 * time.Second)
}
