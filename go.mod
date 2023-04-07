module github.com/breez/breez

require (
	github.com/btcsuite/btcd v0.23.3
	github.com/btcsuite/btcd/btcutil v1.1.1
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/btcwallet v0.15.1
	github.com/btcsuite/btcwallet/walletdb v1.4.0
	github.com/btcsuite/btcwallet/wtxmgr v1.5.0
	github.com/decred/dcrd/lru v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/lightninglabs/neutrino v0.14.2
	github.com/lightninglabs/protobuf-hex-display v1.4.3-hex-display
	github.com/lightningnetwork/lnd v0.15.0-beta
	github.com/status-im/doubleratchet v0.0.0-20181102064121-4dcb6cba284a
	github.com/urfave/cli v1.22.4
	go.etcd.io/bbolt v1.3.6
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/oauth2 v0.0.0-20210615190721-d04028783cf1
	google.golang.org/api v0.30.0
	google.golang.org/grpc v1.38.0
)

replace (
	git.schwanenlied.me/yawning/bsaes.git => github.com/Yawning/bsaes v0.0.0-20180720073208-c0276d75487e
	github.com/btcsuite/btcwallet => github.com/breez/btcwallet v0.15.2-0.20220717090508-739787f948a6
	github.com/btcsuite/btcwallet/walletdb => github.com/breez/btcwallet/walletdb v1.4.1-0.20220717090508-739787f948a6
	github.com/btcsuite/btcwallet/wtxmgr => github.com/breez/btcwallet/wtxmgr v1.5.1-0.20220717090508-739787f948a6
	github.com/lightninglabs/neutrino => github.com/djkazic/neutrino v1.2.3
	github.com/lightningnetwork/lnd => github.com/breez/lnd v0.15.0-beta.rc6.0.20220717110818-cf98538153c2
	github.com/lightningnetwork/lnd/cert => github.com/breez/lnd/cert v1.1.2-0.20220717110818-cf98538153c2
)

go 1.13
