/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-07-22
 */
package p2p

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/peer"
	mplex "github.com/libp2p/go-libp2p-mplex"
	"github.com/libp2p/go-libp2p-secio"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
	"github.com/multiformats/go-multiaddr"
	"v2ray.com/core/app/p2p/account"
	"v2ray.com/core/app/p2p/grpc"
	"v2ray.com/core/app/p2p/protocol"
	"v2ray.com/core/app/p2p/protocol/seedlist"
	"v2ray.com/core/app/p2p/wire"
)

const DefaultSleepInterval = 100 * time.Millisecond

type gRpcService struct {
	NeighborList sync.Map
}

// SayHello implements helloworld.GreeterServer
func (s *gRpcService) SayHello(ctx context.Context, in *wire.HelloSeedList) (*wire.HelloReply, error) {
	fmt.Printf("Received: %v", in.Action)
	s.NeighborList.Store(in.Seed.HostID, in)
	return &wire.HelloReply{}, nil //*wire.HelloReply must not be nil
}

type P2PNode struct {
	gRpcService
	Host      core.Host
	Protocols protocol.Protocol
	ctx       context.Context
	Config    *Config
}

func (this *P2PNode) DoSeedListRequest(pid peer.ID) {
	this.Protocols.RequestSeedList(pid)
}

func (this *P2PNode) onceBootstrap(seed *Seed) peer.ID {
	targetAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/%s/%d/p2p/%s", seed.Ip, seed.Protocol, seed.Port, seed.PubId))
	if err != nil {
		panic(err)
	}

	targetInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		panic(err)
	}

	err = this.Host.Connect(this.ctx, *targetInfo)
	if err != nil {
		panic(err)
	}
	time.Sleep(DefaultSleepInterval)
	return targetInfo.ID
}

func (this *P2PNode) Bootstrap() {
	for _, seed := range this.Config.Seedlist {
		fmt.Println("bootstrap once start:", seed.PubId, seed.Port, seed.Protocol, seed.Ip)
		continue
		this.onceBootstrap(seed)
	}
	wire.SeedlistNotice <- wire.HelloSeedList{
		Action: wire.ActionType_SEED_ONLINE,
		Seed: &wire.SeedInfo{
			Protocol: this.Config.Protocol,
			Ip:       this.Config.Ip,
			Port:     this.Config.Port,
			HostID:   this.Host.ID().Pretty(),
		}}
}

func (this *P2PNode) StartListen(done chan struct{}) {
	p := account.GetAccount().GetPrivKey()

	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)

	listenAddrs := libp2p.ListenAddrStrings(
		fmt.Sprintf("/ip4/%s/%s/%d", this.Config.Ip, this.Config.Protocol, this.Config.Port),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
	)

	security := libp2p.Security(secio.ID, secio.New)
	var err error
	this.Host, err = libp2p.New(
		context.Background(),
		transports,
		muxers,
		security,
		listenAddrs,
		libp2p.Identity(p),
	)

	if err != nil {
		fmt.Printf("libp2p.New in StartListen, err:%s", err.Error())
	}
	seedNode := seedlist.NewSeedNode(this.Host, done)
	this.Protocols.SeedListProtocol = seedlist.NewSeedListProtocol(seedNode, done)
	fmt.Println("host.ID:", this.Host.ID().Pretty())
}

func NewP2PNode(config *Config) *P2PNode {
	return &P2PNode{
		ctx:    context.Background(),
		Config: config}
}

func (this *P2PNode) StartGrpcServer() {
	go grpc.ServiceStart(this)
}

func (this *P2PNode) StartService() {
	this.StartListen(make(chan struct{}))
	this.Protocols.StartSeedlistGossipPubSub(this.ctx, this.Host)
	this.Bootstrap()
	this.StartGrpcServer()
}
