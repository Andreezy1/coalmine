// - Сильный шахтёр:
//     - Оплата труда: 450 угля
//     - Энергия: хватит на 60 добыч угля
//     - Получает угля за одну добычу: 10
//     - Перерыв между добычами: 1 секнда
//     - С каждой новой добычей, характеристика "Получает угля за одну добычу" увеличивается на 3 единицы

package miners

import (
	"context"
	"fmt"
	"time"
)

const (
	PriceforbuyHighMiner int = 450
	timesleepHighMiner       = 1 * time.Second
	energyHighMiner          = 60
	powerHighMiner           = 10
	HighMinerClass           = "Сильный шахтер"
)

var highID = 0

type HighMiner struct {
	Id         int
	Energy     int
	Power      int
	MinerClass string
}

func NewHighMiner() *HighMiner {
	highID++
	return &HighMiner{Id: highID,
		Energy:     energyHighMiner,
		Power:      powerHighMiner,
		MinerClass: HighMinerClass,
	}
}

func (m *HighMiner) Run(ctx context.Context) <-chan Coal {
	coal := make(chan Coal)
	go func() {
		defer close(coal)

		for m.Energy > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(timesleepHighMiner):
			}
			select {
			case <-ctx.Done():
				return
			case coal <- Coal(m.Power):
				m.Power += 3
				m.Energy--
				fmt.Println(m.Power)
			}
		}
	}()
	return coal
}

func (m *HighMiner) Info() MinerInfo {
	return MinerInfo{
		Id:         m.Id,
		Energy:     m.Energy,
		MinerClass: m.MinerClass,
	}
}
