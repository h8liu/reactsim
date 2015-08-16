package sim

import (
	"github.com/h8liu/reactsim/react/sim/structs"
)

type Monitor interface {
	// Tells the demand for the current tick.
	Tell(demand structs.Matrix)
}
