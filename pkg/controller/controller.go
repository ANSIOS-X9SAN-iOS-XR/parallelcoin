package controller

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	blockchain "github.com/p9c/pod/pkg/chain"
	"github.com/p9c/pod/pkg/chain/fork"
	"github.com/p9c/pod/pkg/chain/mining"
	"github.com/p9c/pod/pkg/conte"
	"github.com/p9c/pod/pkg/log"
)

type Blocks []*mining.BlockTemplate

func Run(cx *conte.Xt) (cancel context.CancelFunc) {
	var ctx context.Context
	var lastBlock atomic.Value // Blocks
	ctx, cancel = context.WithCancel(context.Background())
	blockChan := make(chan Blocks)
	go func() {
		for {
			// work loop
			select {
			case lb := <-blockChan:
				lastBlock.Store(lb)
				// send out block broadcast
			case <-ctx.Done():
				// cancel has been called
				return
			default:
			}
		}
	}()
	// create subscriber for new block event
	cx.RPCServer.Cfg.Chain.Subscribe(func(n *blockchain.
	Notification) {
		switch n.Type {
		case blockchain.NTBlockConnected:
			lastBlock.Store(Blocks{})
			// generate Blocks
			for algo := range fork.List[fork.GetCurrent(cx.RPCServer.Cfg.Chain.
				BestSnapshot().Height+1)].Algos {
				// Choose a payment address at random.
				rand.Seed(time.Now().UnixNano())
				payToAddr := cx.StateCfg.ActiveMiningAddrs[rand.Intn(len(cx.
					StateCfg.ActiveMiningAddrs))]
				template, err := cx.RPCServer.Cfg.Generator.NewBlockTemplate(0,
					payToAddr, algo)
				if err != nil {
					log.ERROR("failed to create new block template:", err)
					continue
				}
				lastBlock.Store(append(lastBlock.Load().(Blocks), template))
			}
			blockChan <- lastBlock.Load().(Blocks)
		}
	})
	// goroutine loop checking for connection and sync status
	go func() {
		for {
			time.Sleep(time.Second)
			connCount := cx.RPCServer.Cfg.ConnMgr.ConnectedCount()
			current := cx.RPCServer.Cfg.SyncMgr.IsCurrent()
			if connCount < 1 || !current {
				lastBlock.Store(Blocks{})
				blockChan <- lastBlock.Load().(Blocks)
			}
			select {
			case <-ctx.Done():
				break
			}
		}
	}()
	return
}
