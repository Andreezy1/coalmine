package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet2/company"
	"pet2/company/equipment"
	"pet2/company/miners"
	"time"
)

type HTTPHandlers struct {
	company *company.Company
	srv     http.Server
}

func NewHTTPhandlers(company *company.Company) *HTTPHandlers {
	return &HTTPHandlers{
		company: company,
	}
}

func (h *HTTPHandlers) HandleCreateMiner(w http.ResponseWriter, r *http.Request) {
	var minerDTO MinnerDTO
	if err := json.NewDecoder(r.Body).Decode(&minerDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			TIme:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := minerDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			TIme:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	miner, err := h.company.CreateMiner(minerDTO.MinerClass)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			TIme:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	b, err := json.MarshalIndent(miner, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
}

func (h *HTTPHandlers) InfoPriceMiner(w http.ResponseWriter, r *http.Request) {
	minerPriceDTO := MinerPriceDTO{
		HighMiner:  miners.PriceforbuyHighMiner,
		MidleMiner: miners.PriceforbuyMidleMiner,
		LowMiner:   miners.PriceforbuyLowMiner,
	}
	b, err := json.MarshalIndent(minerPriceDTO, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
}

func (h *HTTPHandlers) AllWorkMiners(w http.ResponseWriter, r *http.Request) {
	allworkminers := h.company.AllInfoMiner()
	info := AllMinersDTO(allworkminers)
	b, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
}

func (h *HTTPHandlers) ClassWorkMiner(w http.ResponseWriter, r *http.Request) {
	class := r.URL.Query().Get("Class")
	classworkminer := h.company.MinerClassInfo(class)
	info := AllMinersDTO(classworkminer)
	b, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}

}

func (h *HTTPHandlers) PriceEquipment(w http.ResponseWriter, r *http.Request) {
	equipmentPriceDTO := EquipmentPriceDTO{
		Vagonetka:    equipment.PriceVagonetka,
		Ventilyaciya: equipment.PriceVentilyaciya,
		Kirka:        equipment.PriceKirka,
	}
	b, err := json.MarshalIndent(equipmentPriceDTO, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
}

func (h *HTTPHandlers) BuyEquipment(w http.ResponseWriter, r *http.Request) {
	var equipmentDTO EquipmentDTO
	if err := json.NewDecoder(r.Body).Decode(&equipmentDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			TIme:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := equipmentDTO.EquipmentValid(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			TIme:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	equipmentbuy, err := h.company.BuyEquipment(equipmentDTO.Equipment)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			TIme:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	b, err := json.MarshalIndent(equipmentbuy, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
}

func (h *HTTPHandlers) CheckBuyEquipment(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(h.company.CheckBuyEquipment(), "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
}

func (h *HTTPHandlers) ALLStatistic(w http.ResponseWriter, r *http.Request) {
	stat := h.company.ALLStatistic()
	statDTO := StatistickDTOconvert(*stat)
	b, err := json.MarshalIndent(statDTO, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
}

func (h *HTTPHandlers) CloseGameHandler(w http.ResponseWriter, r *http.Request) {
	comletedgame, err := h.company.CloseGame()
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			TIme:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	statDTO := StatistickDTOconvert(*comletedgame)
	b, err := json.MarshalIndent(statDTO, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response", err)
		return
	}
	// time.Sleep(5 * time.Second)
	h.company.ContextClose()
}
