package miners

import (
	"context"
	"fmt"
	"time"
)

// - Нормальный шахтёр:
//     - Оплата труда: 50 угля
//     - Энергия: хватит на 45 добыч угля
//     - Получает угля за одну добычу: 3
//     - Перерыв между добычами: 2 секунды
//     - Характеристики с каждой добычей не прогрессируют

const (
	PriceforbuyMidleMiner int = 50
	timesleepMidleMiner       = 2 * time.Second
	energyMidleMiner          = 45
	powerMidleMiner           = 3
	MidleMinerClass           = "Средний шахтер"
)

var midleID = 0

type MidleMiner struct {
	Id         int
	Energy     int
	Power      int
	MinerClass string
}

func NewMidleMiner() *MidleMiner {
	midleID++
	return &MidleMiner{Id: midleID,
		Energy:     energyMidleMiner,
		Power:      powerMidleMiner,
		MinerClass: MidleMinerClass,
	}
}

func (m *MidleMiner) Run(ctx context.Context) <-chan Coal {
	coal := make(chan Coal)
	go func() {
		defer close(coal)
		for m.Energy > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(timesleepMidleMiner):
			}
			select {
			case <-ctx.Done():
				return
			case coal <- Coal(m.Power):
				fmt.Println(m.Power)
				m.Energy--
			}
		}
	}()
	return coal
}

func (m *MidleMiner) Info() MinerInfo {
	return MinerInfo{
		Id:         m.Id,
		Energy:     m.Energy,
		MinerClass: m.MinerClass,
	}
}
