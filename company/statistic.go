package company

import (
	"pet2/company/equipment"
	"pet2/company/miners"
	"sync/atomic"
	"time"
)

type Statistic struct {
	Coal      atomic.Int32
	Miners    map[miners.Miner]miners.Miner
	TimeStart time.Time
	TimeEnd   *time.Time
	Equipment equipment.Equipment
}

func NewStatistic() *Statistic {
	return &Statistic{
		Coal:      atomic.Int32{},
		Miners:    make(map[miners.Miner]miners.Miner),
		TimeStart: time.Now(),
		TimeEnd:   nil,
		Equipment: equipment.Equipment{},
	}
}
