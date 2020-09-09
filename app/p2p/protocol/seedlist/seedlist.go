package seedlist

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/peer"
	protobufCodec "github.com/multiformats/go-multicodec/protobuf"
	"v2ray.com/core/app/p2p/wire"
)

// pattern: /protocol-name/request-or-response-message/version
const seedRequest = "/seedlist/request/0.0.1"
const seedResponse = "/seedlist/response/0.0.1"

// node client version
const ClientVersion = "go-p2p-node/0.0.1"

type Seed struct {
	Hash      string
	Host      string //ip:port
	Timestamp time.Time
}

type SeedListProtocol struct {
	node        *SeedNode                        // local host
	requests    map[string]*wire.SeedListRequest // used to access request data from response handlers
	quit        chan struct{}
	workSeed    []Seed //read from local seedlist file and add from newSeed slic
	newSeed     []Seed // new seed from neighbor node sent
	invalidSeed []Seed // does not work seed
}

// Node type - a p2p host implementing one or more p2p protocols
type SeedNode struct {
	core.Host         // lib-p2p host
	*SeedListProtocol // seedlist protocol impl
}

// Create a new node with its implemented protocols
func NewSeedNode(host core.Host, done chan struct{}) *SeedNode {
	node := &SeedNode{Host: host}
	node.SeedListProtocol = NewSeedListProtocol(node, done)
	return node
}

func NewSeedListProtocol(node *SeedNode, done chan struct{}) *SeedListProtocol {
	p := &SeedListProtocol{node: node, requests: make(map[string]*wire.SeedListRequest), quit: done}
	node.SetStreamHandler(seedRequest, p.onSeedListRequest)
	node.SetStreamHandler(seedResponse, p.onSeedListResponse)
	return p
}

// remote peer requests handler
func (seed *SeedListProtocol) onSeedListRequest(s core.Stream) {

	// get request data
	data := &wire.SeedListRequest{}
	decoder := protobufCodec.Multicodec(nil).Decoder(bufio.NewReader(s))
	err := decoder.Decode(data)
	if err != nil {
		log.Println(err)
		return
	}

	valid := AuthenticateMessage(data, data.MessageData)

	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	// generate response message
	resp := &wire.SeedListResponse{MessageData: NewMessageData(seed.node, data.MessageData.Id, false),
		Message: fmt.Sprintf("Ping response from %s", seed.node.ID())}

	// sign the data
	signature, err := SignProtoMessage(seed.node, resp)
	if err != nil {
		log.Println("failed to sign response")
		return
	}

	// add the signature to the message
	resp.MessageData.Sign = signature

	// send the response
	s, respErr := seed.node.NewStream(context.Background(), s.Conn().RemotePeer(), seedResponse)
	if respErr != nil {
		log.Println(respErr)
		return
	}

	ok := SendProtoMessage(resp, s)

	if !ok {
		log.Printf("%s: SeedList response to %s sent error.", s.Conn().LocalPeer().String(), s.Conn().RemotePeer().String())
	} else {
		log.Printf("%s: SeedList response to %s sent success, response:%s", s.Conn().LocalPeer().String(), s.Conn().RemotePeer().String(), resp.Message)

	}
}

// remote ping response handler
func (seed *SeedListProtocol) onSeedListResponse(s core.Stream) {
	data := &wire.SeedListResponse{}
	decoder := protobufCodec.Multicodec(nil).Decoder(bufio.NewReader(s))
	err := decoder.Decode(data)
	if err != nil {
		return
	}

	valid := AuthenticateMessage(data, data.MessageData)

	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	// locate request data and remove it if found
	_, ok := seed.requests[data.MessageData.Id]
	if ok {
		// remove request from map as we have processed it here
		delete(seed.requests, data.MessageData.Id)
	} else {
		log.Println("Failed to locate request data boject for response")
		return
	}

	log.Printf("%s: Received seedlist response from %s. Message id:%s. Message: %s.", s.Conn().LocalPeer(), s.Conn().RemotePeer(), data.MessageData.Id, data.Message)
	close(seed.quit)
}

func (seed *SeedListProtocol) RequestSeedList(pid peer.ID) bool {
	log.Printf("%s: Sending seedlist request to: %s....", seed.node.ID(), pid)

	// create message data
	req := &wire.SeedListRequest{MessageData: NewMessageData(seed.node, uuid.New().String(), false),
		Message: fmt.Sprintf("Request Seedlist from %s", seed.node.ID())}

	// sign the data
	signature, err := SignProtoMessage(seed.node, req)
	if err != nil {
		log.Println("failed to sign pb data")
		return false
	}

	// add the signature to the message
	req.MessageData.Sign = signature

	s, err := seed.node.NewStream(context.Background(), pid, seedRequest)
	if err != nil {
		log.Println(err)
		return false
	}

	ok := SendProtoMessage(req, s)

	if !ok {
		return false
	}

	// store ref request so response handler has access to it
	seed.requests[req.MessageData.Id] = req
	log.Printf("%s: Request SeedList from: %s was sent. Message Id: %s, Message: %s", seed.node.ID(), pid, req.MessageData.Id, req.Message)
	return true
}
