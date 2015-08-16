package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h8liu/reactsim/react/conf"
	dem "github.com/h8liu/reactsim/react/demand"
	"github.com/h8liu/reactsim/react/sim"
	. "github.com/h8liu/reactsim/react/sim/config"
	"github.com/h8liu/reactsim/react/sim/demand"
	"github.com/h8liu/reactsim/react/sim/diag"
	"github.com/h8liu/reactsim/react/sim/tors"
)

type runConfig struct {
	NotHybrid        bool
	Eighty           bool
	LargeFlowPercent int
	Nsmall           int
}

type runResult struct {
	GoodputPercent float64
}

func runSetup(c *runConfig) {
	conf.Load()

	SetNhost(64)
	TickPerGrid = 1
	LinkBw = 12500
	if c.Eighty {
		PackBw = 2500
	} else {
		PackBw = 1250
	}

	if c.NotHybrid {
		PackBw = 0
	}

	NightLen = 20
	MinDayLen = 40
	AvgDayLen = 500
	WeekLen = 3000
	NicBufCover = 5000

	Tracking = false
}

func makeDemand(c *runConfig, t uint64) *dem.Demand {
	ret := dem.NewDemand(Nhost)
	p := dem.NewPeriod(Nhost, t)

	bigBw := LinkBw * uint64(c.LargeFlowPercent) / 100
	smallBw := (LinkBw - bigBw) / uint64(c.Nsmall)

	for i := 0; i < Nhost; i++ {
		p.D[i][(i+1)%Nhost] = bigBw
		for j := 0; j < c.Nsmall; j++ {
			p.D[i][(i+2+j)%Nhost] = smallBw
		}
	}

	ret.Add(p)
	return ret
}

func makeScheduler(c *runConfig) sim.Scheduler {
	s := diag.NewScheduler()
	s.SafeBandw = false
	s.PureCircuit = c.NotHybrid
	return s
}

func run(c *runConfig) *runResult {
	runSetup(c)
	t := WeekLen * 3
	d := makeDemand(c, t)
	hosts := demand.NewHosts(d)
	hosts.Restart()
	sw := tors.NewReactSwitch()
	sched := makeScheduler(c)
	testbed := sim.NewTestbed(hosts, sw)
	testbed.Estimator = hosts
	testbed.Scheduler = sched
	testbed.Progress = sim.NewLineProgress()
	testbed.WarmUp = WeekLen

	if err := testbed.Run(t); err != nil {
		panic(err)
	}

	goodput := testbed.Goodput
	capacity := testbed.Capacity

	return &runResult{float64(goodput) / float64(capacity) * 100}
}

func createFile(fout string) *os.File {
	if fout == "" {
		return nil
	}
	f, err := os.Create(fout)
	if err != nil {
		panic(err)
	}
	return f
}

func runFig10(fout string) {
	f := createFile(fout)
	if f != nil {
		defer f.Close()
	}

	for largep := 100; largep >= 10; largep-- {
		resCirc := run(&runConfig{
			NotHybrid:        true,
			LargeFlowPercent: largep,
			Nsmall:           20,
		})

		resHybrid20 := run(&runConfig{
			LargeFlowPercent: largep,
			Nsmall:           20,
			Eighty:           true,
		})

		resHybrid10 := run(&runConfig{
			LargeFlowPercent: largep,
			Nsmall:           20,
		})

		fmt.Printf("large=%d%% circ=%.3f%% hybr10=%.3f%% hybr20=%.3f%%\n",
			largep,
			resCirc.GoodputPercent,
			resHybrid10.GoodputPercent,
			resHybrid20.GoodputPercent,
		)

		if f != nil {
			fmt.Fprintf(f, "%d\t%.3f\t%.3f\t%.3f\n",
				largep,
				resCirc.GoodputPercent,
				resHybrid10.GoodputPercent,
				resHybrid20.GoodputPercent,
			)
		}
	}
}

func runFig11(fout string) {
	f := createFile(fout)
	if f != nil {
		defer f.Close()
	}

	for nsmall := 1; nsmall <= 62; nsmall++ {
		resCirc := run(&runConfig{
			NotHybrid:        true,
			LargeFlowPercent: 90,
			Nsmall:           nsmall,
		})
		resHybrid := run(&runConfig{
			LargeFlowPercent: 90,
			Nsmall:           nsmall,
		})

		fmt.Printf("nsmall=%d circ=%.3f%% hybr=%.3f%%\n",
			nsmall,
			resCirc.GoodputPercent,
			resHybrid.GoodputPercent,
		)

		if f != nil {
			fmt.Fprintf(f, "%d\t%.3f\t%.3f\n",
				nsmall,
				resCirc.GoodputPercent,
				resHybrid.GoodputPercent,
			)
		}
	}
}

var (
	fig10Out = flag.String("fig10", "", "path to save data of figure10")
	fig11Out = flag.String("fig11", "", "path to save data of figure11")
)

func main() {
	flag.Parse()

	fmt.Println("[fig10]")
	runFig10(*fig10Out)

	fmt.Println("[fig11]")
	runFig11(*fig11Out)
}
