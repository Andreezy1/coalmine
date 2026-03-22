package miners

import (
	"context"
	"fmt"
	"time"
)

// - Маленький шахтёр:
//     - Оплата труда: 5 угля
//     - Энергии хватит на: 30 добыч угля
//     - Получает за одну добычу: 1 уголь
//     - Перерыв между добычами: 3 секунды
//     - Характеристики с каждой добычей не прогрессируют

const (
	PriceforbuyLowMiner int = 5
	timesleepLowMiner       = 3 * time.Second
	energyLowMiner          = 30
	powerLowMiner           = 1
	LowMinerClass           = "Младший шахтер"
)

var lowID = 0

type LowMiner struct {
	Id         int
	Energy     int
	Power      int
	MinerClass string
}

func NewLowMiner() *LowMiner {
	lowID++
	return &LowMiner{Id: lowID,
		Energy:     energyLowMiner,
		Power:      powerLowMiner,
		MinerClass: LowMinerClass,
	}
}

func (m *LowMiner) Run(ctx context.Context) <-chan Coal {
	coal := make(chan Coal)
	go func() {
		defer close(coal)

		for m.Energy > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(timesleepLowMiner):
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

func (m *LowMiner) Info() MinerInfo {
	return MinerInfo{
		Id:         m.Id,
		Energy:     m.Energy,
		MinerClass: m.MinerClass,
	}
}
