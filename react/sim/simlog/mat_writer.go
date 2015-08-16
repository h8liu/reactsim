package simlog

import (
	"encoding/json"
	"os"

	"github.com/h8liu/reactsim/react/sim/clock"
	"github.com/h8liu/reactsim/react/sim/structs"
)

type MatWriter struct {
	fout *os.File
	enc  *json.Encoder
}

func ne(e error) {
	if e != nil {
		panic(e)
	}
}

func NewMatWriter(path string) *MatWriter {
	fout, e := os.Create(path)
	ne(e)

	ret := new(MatWriter)
	ret.fout = fout
	ret.enc = json.NewEncoder(fout)
	return ret
}

func (self *MatWriter) Write(m structs.Matrix) {
	entry := &MatEntry{
		clock.T,
		m,
	}

	ne(self.enc.Encode(entry))
}

func (self *MatWriter) Close() {
	ne(self.fout.Close())
}
