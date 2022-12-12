package discoveryimpl

import (
	"github.com/xiaomdong/go-fil-markets/discovery"
)

func Multi(r discovery.PeerResolver) discovery.PeerResolver { // TODO: actually support multiple mechanisms
	return r
}
