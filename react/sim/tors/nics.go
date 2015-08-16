package tors

import (
	. "github.com/h8liu/reactsim/react/sim/config"
	. "github.com/h8liu/reactsim/react/sim/queues"
)

func NewNics() *Queues {
	return NewSizedQueues(NicBufSize())
	// return NewNicQueues(NicBufSize(), NicBufTotalSize())
}
