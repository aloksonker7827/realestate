package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gitlab.com/umi7/DezHab_user_backend/api"
	"gitlab.com/umi7/DezHab_user_backend/api/handler"
	"gitlab.com/umi7/DezHab_user_backend/api/rest"
	"gitlab.com/umi7/DezHab_user_backend/config"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"gitlab.com/umi7/DezHab_user_backend/utils"
	"net/http"
	"os"
	"time"
)

const SwaggerPath = "/src/gitlab.com/umi7/DezHab_user_backend/resources/swaggerui/"

// Server represents a server mux
type Server struct {
	*mux.Router
	Address string
}

// New setups & returns a server
func New() *Server {
	router := mux.NewRouter()
	address := config.AppConfig.Server.IP + ":" + config.AppConfig.Server.Port
	s := Server{Router: router, Address: address}
	s.SetupRouter()
	return &s
}

// SetupRouters configures the route for the server
func (s Server) SetupRouter() {
	s.HandleFunc("/status", api.HealthCheck).Methods(http.MethodGet)
	// For swagger
	dirPath := utils.GoPath + SwaggerPath
	swaggerFileHandler := http.FileServer(http.Dir(dirPath))
	swaggerHandler := http.StripPrefix("/swaggerui/", swaggerFileHandler)
	s.PathPrefix("/swaggerui/").Handler(swaggerHandler)
	apiMux := s.PathPrefix("/api").Subrouter()
	v1Router := apiMux.PathPrefix("/v1").Subrouter()
	v2Router := apiMux.PathPrefix("/v2").Subrouter()
	s.registerV1Route(v1Router)
	s.registerV2Route(v2Router)
}
func (s Server) registerV1Route(router *mux.Router) {
	router.HandleFunc("/status", handler.MakeRest(dto.GetBookKeepingRequest{}, rest.GetBookKeeping)).Methods(http.MethodGet)
	router.HandleFunc("/status", handler.MakeRest(dto.UpdateBookKeepingRequest{}, rest.UpdateBookKeeping)).Methods(http.MethodPut)
}

func (s Server) registerV2Route(router *mux.Router) {
}

func (s Server) ServeHTTP() {
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.Router)
	srv := &http.Server{
		Handler: loggedRouter,
		Addr:    s.Address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}
	logrus.Fatal(srv.ListenAndServe())
}
