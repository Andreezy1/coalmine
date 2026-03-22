package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/Mainers").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateMiner)
	router.Path("/Mainers").Methods("GET").Queries("Class", "").HandlerFunc(s.httpHandlers.ClassWorkMiner)
	router.Path("/Mainers").Methods("GET").HandlerFunc(s.httpHandlers.AllWorkMiners)
	router.Path("/Mainers/Price").Methods("GET").HandlerFunc(s.httpHandlers.InfoPriceMiner)

	router.Path("/Equipment").Methods("POST").HandlerFunc(s.httpHandlers.BuyEquipment)
	router.Path("/Equipment").Methods("GET").HandlerFunc(s.httpHandlers.CheckBuyEquipment)
	router.Path("/Equipment/Price").Methods("GET").HandlerFunc(s.httpHandlers.PriceEquipment)

	router.Path("/Company").Methods("GET").HandlerFunc(s.httpHandlers.ALLStatistic)
	router.Path("/Company/Close").Methods("POST").HandlerFunc(s.httpHandlers.CloseGameHandler)

	srv := http.Server{
		Addr:    ":9091",
		Handler: router,
	}

	go func() {
		<-s.httpHandlers.company.MinerContex.Done()

		srv.Shutdown(s.httpHandlers.company.MinerContex)
	}()

	if err := srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil

}
