package company

import (
	"context"
	"fmt"
	"pet2/company/equipment"
	"pet2/company/miners"
	"sync"
	"time"
)

type Company struct {
	Transferchanelcoal chan miners.Coal
	Mtx                sync.Mutex
	// Miners             map[int]miners.Miner
	// Time               time.Time
	// Coal               atomic.Int32
	// Equipment          equipment.Equipment
	MinerContex      context.Context
	ContextClose     context.CancelFunc
	StatisticCompany Statistic
}

func NewCompany() *Company {
	minerContext, contextCLose := context.WithCancel(context.Background())
	c := &Company{
		Transferchanelcoal: make(chan miners.Coal),
		Mtx:                sync.Mutex{},
		// Miners:             make(map[int]miners.Miner),
		// Time:               time.Now(),
		// Coal:               atomic.Int32{},
		// Equipment:          equipment.Equipment{},
		MinerContex:      minerContext,
		ContextClose:     contextCLose,
		StatisticCompany: *NewStatistic(),
	}
	go c.PlusCoal()
	go c.ChanTransferCoal()
	return c
}

func (c *Company) PlusCoal() {
	// c.StatisticCompany.Coal.Add(1000)
	for {
		time.Sleep(1 * time.Second)
		c.StatisticCompany.Coal.Add(1)
	}
}

func (c *Company) ChanTransferCoal() {
	for v := range c.Transferchanelcoal {
		c.StatisticCompany.Coal.Add(int32(v))
	}
}

// var i int = 0

func (c *Company) CreateMiner(typeMiner string) (miners.Miner, error) {
	var miner miners.Miner
	c.Mtx.Lock()
	defer c.Mtx.Unlock()
	switch typeMiner {
	case miners.LowMinerClass:
		if c.StatisticCompany.Coal.Load() < int32(miners.PriceforbuyLowMiner) {
			return miner, ErrorNotMoney
		}
		c.StatisticCompany.Coal.Add(-int32(miners.PriceforbuyLowMiner))
		fmt.Println(c.StatisticCompany.Coal.Load(), miners.PriceforbuyLowMiner)
		miner = miners.NewLowMiner()
	case miners.MidleMinerClass:
		if c.StatisticCompany.Coal.Load() < int32(miners.PriceforbuyMidleMiner) {
			return miner, ErrorNotMoney
		}
		c.StatisticCompany.Coal.Add(-int32(miners.PriceforbuyMidleMiner))
		fmt.Println(c.StatisticCompany.Coal.Load(), miners.PriceforbuyMidleMiner)
		miner = miners.NewMidleMiner()
	case miners.HighMinerClass:
		if c.StatisticCompany.Coal.Load() < int32(miners.PriceforbuyHighMiner) {
			return miner, ErrorNotMoney
		}
		c.StatisticCompany.Coal.Add(-int32(miners.PriceforbuyHighMiner))
		fmt.Println(c.StatisticCompany.Coal.Load(), miners.PriceforbuyHighMiner)
		miner = miners.NewHighMiner()
	default:
		return miner, nil
	}
	trpoint := miner.Run(c.MinerContex)

	c.StatisticCompany.Miners[miner] = miner
	go func() {
		for v := range trpoint {
			c.Transferchanelcoal <- v
		}
	}()
	return miner, nil
}

func (c *Company) InfoCoal() {
	fmt.Println(c.StatisticCompany.Coal.Load())
}

func (c *Company) AllInfoMiner() map[miners.Miner]miners.MinerInfo {
	tmp := make(map[miners.Miner]miners.MinerInfo)
	for k, v := range c.StatisticCompany.Miners {
		tmp[k] = v.Info()
	}
	return tmp
}

func (c *Company) MinerClassInfo(mci string) map[miners.Miner]miners.MinerInfo {
	tmp := make(map[miners.Miner]miners.MinerInfo)
	for k, v := range c.StatisticCompany.Miners {
		if mci == v.Info().MinerClass {
			tmp[k] = v.Info()
		}
	}
	return tmp
}

func (c *Company) BuyEquipment(equip string) (equipment.Equipment, error) {
	c.Mtx.Lock()
	defer c.Mtx.Unlock()
	switch equip {
	case "Кирка":
		if c.StatisticCompany.Equipment.Kirka {
			return c.StatisticCompany.Equipment, ErrorHaveItem
		}
		if c.StatisticCompany.Coal.Load() < equipment.PriceKirka {
			return c.StatisticCompany.Equipment, ErrorNotMoney
		}
		c.StatisticCompany.Coal.Add(-equipment.PriceKirka)
		c.StatisticCompany.Equipment.BuyKirka()
		fmt.Println(c.StatisticCompany.Equipment)
	case "Вентиляция":
		if c.StatisticCompany.Equipment.Ventilyaciya {
			return c.StatisticCompany.Equipment, ErrorHaveItem
		}
		if c.StatisticCompany.Coal.Load() < equipment.PriceVentilyaciya {
			return c.StatisticCompany.Equipment, ErrorNotMoney
		}
		c.StatisticCompany.Coal.Add(-equipment.PriceVentilyaciya)
		c.StatisticCompany.Equipment.BuyVentilyaciya()
		fmt.Println(c.StatisticCompany.Equipment)
	case "Вагонетка":
		if c.StatisticCompany.Equipment.Vagonetka {
			return c.StatisticCompany.Equipment, ErrorHaveItem
		}
		if c.StatisticCompany.Coal.Load() < equipment.PriceVagonetka {
			return c.StatisticCompany.Equipment, ErrorNotMoney
		}
		c.StatisticCompany.Coal.Add(-equipment.PriceVagonetka)
		c.StatisticCompany.Equipment.BuyVagonetka()
		fmt.Println(c.StatisticCompany.Equipment)
	}
	return c.StatisticCompany.Equipment, nil
}

func (c *Company) CheckBuyEquipment() equipment.Equipment {
	return c.StatisticCompany.Equipment
}

func (c *Company) ALLStatistic() *Statistic {
	return &c.StatisticCompany
}

func (c *Company) CloseGame() (*Statistic, error) {
	if c.StatisticCompany.Equipment.Kirka && c.StatisticCompany.Equipment.Vagonetka && c.StatisticCompany.Equipment.Ventilyaciya {
		// c.ContextClose()
		time := time.Now()
		c.StatisticCompany.TimeEnd = &time
		return c.ALLStatistic(), nil
	}
	return c.ALLStatistic(), ErrorCloseGame
}
