package server

import (
	"encoding/json"
	"errors"
	"pet2/company"
	"pet2/company/equipment"
	"pet2/company/miners"
	"time"
)

type MinnerDTO struct {
	MinerClass string `json: MinerClass`
}

func (m *MinnerDTO) ValidateForCreate() error {
	if m.MinerClass == "" {
		return errors.New("Miner is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	TIme    time.Time
}

func CreateErrDTO(message string) *ErrorDTO {
	return &ErrorDTO{
		Message: message,
		TIme:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

type MinerPriceDTO struct {
	HighMiner  int
	MidleMiner int
	LowMiner   int
}

type EquipmentPriceDTO struct {
	Vagonetka    int
	Ventilyaciya int
	Kirka        int
}

type EquipmentDTO struct {
	Equipment string `json: Equipment`
}

func (e *EquipmentDTO) EquipmentValid() error {
	if e.Equipment == "" {
		return errors.New("Equipment is empty")
	}
	return nil
}

func AllMinersDTO(m map[miners.Miner]miners.MinerInfo) []miners.MinerInfo {
	var info []miners.MinerInfo
	for _, v := range m {
		info = append(info, v)
	}
	return info
}

type StatisticDTO struct {
	Coal      int
	Miners    []miners.MinerInfo
	TimeStart time.Time
	TimeEnd   *time.Time
	Equipment equipment.Equipment
}

func StatistickDTOconvert(st company.Statistic) StatisticDTO {
	var tmp []miners.MinerInfo
	for _, v := range st.Miners {
		tmp = append(tmp, v.Info())
	}
	return StatisticDTO{
		Coal:      int(st.Coal.Load()),
		Miners:    tmp,
		TimeStart: st.TimeStart,
		TimeEnd:   st.TimeEnd,
		Equipment: st.Equipment,
	}
}
