package k6setenv

import "github.com/go-ping/ping"

type LgcPinger struct {
	Pinger *ping.Pinger
	Err    error
	Stats  *ping.Statistics
}

func (lp *LgcPinger) SetNewPinger(url string) {
	lp.Pinger, lp.Err = ping.NewPinger(url)
}

func (lp *LgcPinger) RunPinger(count int) {
	lp.Pinger.Count = count
	lp.Err = lp.Pinger.Run()

	if lp.Err != nil {
		panic(lp.Err)
	}
}

func (lp *LgcPinger) GetStats() *ping.Statistics {
	lp.Stats = lp.Pinger.Statistics()
	return lp.Stats
}
