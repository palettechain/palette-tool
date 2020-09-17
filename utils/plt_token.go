package utils

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-log"
)

var (
	logger     = log.Logger("palette")
	PLTAddress = common.HexToAddress(native.PLTContractAddress)
)

func init() {
	plt.InitABI()
}
//
//func (c *PaletteClient) PLTMint(amount *big.Int) (common.Hash, error) {
//	payload, err := utils.PackMethod(plt.ABI, plt.MethodMint, c.AdminAddress(), amount)
//	if err != nil {
//		return common.Hash{}, err
//	}
//
//	return c.SendPLTTransaction(c.Admin.PrivateKey, payload)
//}

func (c *PaletteClient) PLTTransfer(key *ecdsa.PrivateKey, to common.Address, amount *big.Int) (common.Hash, error) {
	payload, err := utils.PackMethod(plt.ABI, plt.MethodTransfer, to, amount)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendPLTTransaction(key, payload)
}

func (c *PaletteClient) PLTTransferFrom(key *ecdsa.PrivateKey, from, to common.Address, amount *big.Int) (common.Hash, error) {
	payload, err := utils.PackMethod(plt.ABI, plt.MethodTransferFrom, from, to, amount)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendPLTTransaction(key, payload)
}

func (c *PaletteClient) PLTApprove(key *ecdsa.PrivateKey, spender common.Address, amount *big.Int) (common.Hash, error) {
	payload, err := utils.PackMethod(plt.ABI, plt.MethodApprove, spender, amount)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendPLTTransaction(key, payload)
}

func (c *PaletteClient) PLTTotalSupply() (*big.Int, error) {
	payload, err := utils.PackMethod(plt.ABI, plt.MethodTotalSupply)
	if err != nil {
		return nil, err
	}

	raw, err := c.CallContract(c.AdminAddress(), PLTAddress, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get total supply: [%v]", err)
	}

	supply := new(big.Int).SetBytes(raw)
	return supply, nil
}

func (c *PaletteClient) PLTDecimals() (uint64, error) {
	payload, err := utils.PackMethod(plt.ABI, plt.MethodDecimals)
	if err != nil {
		return 0, err
	}

	raw, err := c.CallContract(c.AdminAddress(), PLTAddress, payload)
	if err != nil {
		return 0, fmt.Errorf("failed to get decimal: [%v]", err)
	}

	decimal := new(big.Int).SetBytes(raw).Uint64()
	return decimal, nil
}

func (c *PaletteClient) SendPLTTransaction(key *ecdsa.PrivateKey, payload []byte) (common.Hash, error) {
	addr := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := c.GetNonce(addr.Hex())
	if err != nil {
		return common.Hash{}, err
	}
	tx := types.NewTransaction(
		nonce,
		PLTAddress,
		big.NewInt(0),
		GasNormal,
		big.NewInt(GasPrice),
		payload,
	)

	signedTx, err := c.SignTransaction(key, tx)
	if err != nil {
		return common.Hash{}, err
	}
	return c.SendRawTransaction(signedTx)
}
