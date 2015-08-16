package queues

import (
	. "github.com/h8liu/reactsim/react/sim/structs"
)

type Puller interface {
	Pull(budget Matrix) Matrix
}
