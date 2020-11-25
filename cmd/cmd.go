// Copyright 2017 AMIS Technologies
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/discv5"
	istcommon "github.com/palettechain/palette-tool/common"
	"github.com/palettechain/palette-tool/genesis"
	"github.com/urfave/cli"
)

type validatorInfo struct {
	Address  common.Address
	Nodekey  string
	NodeInfo string
}

var (
	SetupCommand = cli.Command{
		Name:  "setup",
		Usage: "Setup your Istanbul network in seconds",
		Description: `This tool helps generate:

		* Genesis block
		* Static nodes for all validators
		* Validator details

	    for Istanbul consensus.
`,
		Action: gen,
		Flags: []cli.Flag{
			numOfValidatorsFlag,
			staticNodesFlag,
			verboseFlag,
			quorumFlag,
			saveFlag,
		},
	}
)

func gen(ctx *cli.Context) error {
	num := ctx.Int(numOfValidatorsFlag.Name)

	keys, nodekeys, addrs := istcommon.GenerateKeys(num)
	var nodes []string

	if ctx.Bool(verboseFlag.Name) {
		fmt.Println("validators")
	}

	if ctx.Bool(saveFlag.Name) {
		os.MkdirAll("setup", os.ModePerm)
	}

	for i := 0; i < num; i++ {
		v := &validatorInfo{
			Address: addrs[i],
			Nodekey: nodekeys[i],
			NodeInfo: discv5.NewNode(
				discv5.PubkeyID(&keys[i].PublicKey),
				net.ParseIP("0.0.0.0"),
				0,
				uint16(30300)).String(),
		}

		nodes = append(nodes, string(v.NodeInfo))

		if ctx.Bool(verboseFlag.Name) {
			str, _ := json.MarshalIndent(v, "", "\t")
			fmt.Println(string(str))

			if ctx.Bool(saveFlag.Name) {
				folderName := fmt.Sprintf("node%d", i) //strconv.Itoa(i)
				folderPath := path.Join("setup", "nodekeys", folderName)
				os.MkdirAll(folderPath, os.ModePerm)
				ioutil.WriteFile(path.Join(folderPath, "nodekey"), []byte(nodekeys[i]), os.ModePerm)
			}
		}
	}

	if ctx.Bool(verboseFlag.Name) {
		fmt.Print("\n\n\n")
	}

	staticNodes, _ := json.MarshalIndent(nodes, "", "\t")
	if ctx.Bool(staticNodesFlag.Name) {
		name := "static-nodes.json"
		fmt.Println(name)
		fmt.Println(string(staticNodes))
		fmt.Print("\n\n\n")

		if ctx.Bool(saveFlag.Name) {
			ioutil.WriteFile(path.Join("setup", name), staticNodes, os.ModePerm)
		}
	}

	var jsonBytes []byte
	isQuorum := ctx.Bool(quorumFlag.Name)
	g := genesis.New(
		genesis.Validators(addrs...),
	)

	if isQuorum {
		jsonBytes, _ = json.MarshalIndent(genesis.ToQuorum(g, true), "", "    ")
	} else {
		jsonBytes, _ = json.MarshalIndent(g, "", "    ")
	}
	fmt.Println("genesis.json")
	fmt.Println(string(jsonBytes))

	if ctx.Bool(saveFlag.Name) {
		ioutil.WriteFile(path.Join("setup", "genesis.json"), jsonBytes, os.ModePerm)
	}

	return nil
}

func removeSpacesAndLines(b []byte) string {
	out := string(b)
	out = strings.Replace(out, " ", "", -1)
	out = strings.Replace(out, "\t", "", -1)
	out = strings.Replace(out, "\n", "", -1)
	return out
}
