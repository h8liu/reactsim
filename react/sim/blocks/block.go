package blocks

import (
	. "github.com/h8liu/reactsim/react/sim/packet"
)

// interface for simulation blocks
type Block interface {
	// unlike send() method in sender, a push method
	// always succeed with all the bytes moved in
	// to the block
	Push(packet *Packet)
}
