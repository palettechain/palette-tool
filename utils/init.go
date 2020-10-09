package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/governance"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/ipfs/go-log"
)

var (
	PLTABI, GovernanceABI abi.ABI
	logger                = log.Logger("palette")
	PLTAddress            = common.HexToAddress(native.PLTContractAddress)
	GovernanceAddress     = common.HexToAddress(native.GovernanceContractAddress)
)

func init() {
	PLTABI = plt.GetABI()
	GovernanceABI = governance.GetABI()
}
