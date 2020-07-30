package lnnode

import (
	"sync"
	"time"

	"github.com/breez/breez/config"
	"github.com/breez/breez/db"
	breezlog "github.com/breez/breez/log"
	"github.com/btcsuite/btclog"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/breezbackuprpc"
	"github.com/lightningnetwork/lnd/lnrpc/chainrpc"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/lightningnetwork/lnd/lnrpc/signrpc"
	"github.com/lightningnetwork/lnd/lnrpc/submarineswaprpc"
	"github.com/lightningnetwork/lnd/lnrpc/walletrpc"
	"github.com/lightningnetwork/lnd/subscribe"
)

// API represents the lnnode exposed functions that are accessible for
// breez services to use.
// It is mainly enable the service to subscribe to various daemon events
// and get an APIClient to query the daemon directly via RPC.
type API interface {
	SubscribeEvents() (*subscribe.Client, error)
	HasActiveChannel() bool
	IsReadyForPayment() bool
	WaitReadyForPayment(timeout time.Duration) error
	NodePubkey() string
	APIClient() lnrpc.LightningClient
	SubSwapClient() submarineswaprpc.SubmarineSwapperClient
	BreezBackupClient() breezbackuprpc.BreezBackuperClient
	RouterClient() routerrpc.RouterClient
	WalletKitClient() walletrpc.WalletKitClient
	ChainNotifierClient() chainrpc.ChainNotifierClient
	InvoicesClient() invoicesrpc.InvoicesClient
	SignerClient() signrpc.SignerClient
}

// Daemon contains data regarding the lightning daemon.
type Daemon struct {
	sync.Mutex
	cfg                 *config.Config
	breezDB             *db.DB
	started             int32
	stopped             int32
	startTime           time.Time
	daemonRunning       bool
	nodePubkey          string
	wg                  sync.WaitGroup
	log                 btclog.Logger
	lightningClient     lnrpc.LightningClient
	subswapClient       submarineswaprpc.SubmarineSwapperClient
	breezBackupClient   breezbackuprpc.BreezBackuperClient
	routerClient        routerrpc.RouterClient
	walletKitClient     walletrpc.WalletKitClient
	chainNotifierClient chainrpc.ChainNotifierClient
	invoicesClient      invoicesrpc.InvoicesClient
	signerClient        signrpc.SignerClient
	ntfnServer          *subscribe.Server
	quitChan            chan struct{}
	startBeforeSync     bool
}

// NewDaemon is used to create a new daemon that wraps a lightning
// network daemon.
func NewDaemon(cfg *config.Config, db *db.DB, startBeforeSync bool) (*Daemon, error) {
	logger, err := breezlog.GetLogger(cfg.WorkingDir, "DAEM")
	if err != nil {
		return nil, err
	}

	return &Daemon{
		cfg:             cfg,
		breezDB:         db,
		ntfnServer:      subscribe.NewServer(),
		log:             logger,
		startBeforeSync: startBeforeSync,
	}, nil
}
