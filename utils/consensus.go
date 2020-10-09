package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/governance"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// params are validator and isRevoke
func (c *PaletteClient) AddValidator(validator common.Address) (common.Hash, error) {
	payload, err := utils.PackMethod(PLTABI, governance.MethodAddValidator, validator, false)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendGovernanceTransaction(c.Admin.PrivateKey, payload)
}

func (c *PaletteClient) SendGovernanceTransaction(key *ecdsa.PrivateKey, payload []byte) (common.Hash, error) {
	addr := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := c.GetNonce(addr.Hex())
	if err != nil {
		return common.Hash{}, err
	}
	tx := types.NewTransaction(
		nonce,
		GovernanceAddress,
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
