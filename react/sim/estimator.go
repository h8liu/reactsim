package sim

import (
	"github.com/h8liu/reactsim/react/sim/structs"
)

type Estimator interface {
	// Returns the demand for the next weeklen
	Estimate() (structs.Matrix, uint64)
}
