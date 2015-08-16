package blocks

import (
	. "github.com/h8liu/reactsim/react/sim/packet"
)

// interface with back pressure
// this is the input interface for a switch
type Sender interface {
	Send(packet *Packet) uint64
}
