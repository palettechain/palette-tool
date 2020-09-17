module github.com/palettechain/palette-tool

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/ethereum/go-ethereum v1.0.0
	github.com/ipfs/go-log v1.0.4
)

replace (
	github.com/coreos/etcd v0.0.1 => github.com/polynetwork/coreos-etcd v0.0.1
	github.com/coreos/go-semver v0.0.1 => github.com/polynetwork/coreos-semver v0.0.1
	github.com/coreos/go-systemd v0.0.1 => github.com/polynetwork/coreos-systemd v0.0.1
	github.com/coreos/pkg v0.0.1 => github.com/polynetwork/coreos-pkg v0.0.1
	github.com/ethereum/go-ethereum v1.0.0 => ../palette
)
