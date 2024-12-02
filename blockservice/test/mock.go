package bstest

import (
	delay "github.com/ipfs/go-ipfs-delay"
	testinstance "github.com/stateless-minds/boxo/bitswap/testinstance"
	tn "github.com/stateless-minds/boxo/bitswap/testnet"
	"github.com/stateless-minds/boxo/blockservice"
	mockrouting "github.com/stateless-minds/boxo/routing/mock"
)

// Mocks returns |n| connected mock Blockservices
func Mocks(n int, opts ...blockservice.Option) []blockservice.BlockService {
	net := tn.VirtualNetwork(delay.Fixed(0))
	routing := mockrouting.NewServer()
	sg := testinstance.NewTestInstanceGenerator(net, routing, nil, nil)
	instances := sg.Instances(n)

	var servs []blockservice.BlockService
	for _, i := range instances {
		servs = append(servs, blockservice.New(i.Blockstore,
			i.Exchange, opts...))
	}
	return servs
}
