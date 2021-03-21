package rest_server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/microdimmer/key-value_rest/tree/main/internal/app/kv_db"
	"github.com/sirupsen/logrus"
)

type RESTServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	db     *DataMap
}

func New(config *Config) *RESTServer {
	return &RESTServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		db:     kv_db.Create(),
	}
}

func (s *RESTServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *RESTServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *RESTServer) configureRouter() {
	// s.router.HandleFunc("/Upsert", s.handleGet())
	// s.router.HandleFunc("/Delete", s.handleGet())
	// s.router.HandleFunc("/get", s.handleGet())
	// s.router.HandleFunc("/list", s.handleGet())
	// s.router.HandleFunc("/hello", s.handleHello())

	s.router.HandleFunc("/list", s.handleList()).Methods("GET")
	s.router.HandleFunc("/get", s.handleGet()).Methods("GET")
	s.router.HandleFunc("/delete", s.handleDelete()).Methods("DELETE")
	s.router.HandleFunc("/upsert", s.handleUpsert()).Methods("POST")

	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *RESTServer) handleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *RESTServer) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *RESTServer) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// s.db.Delete()
		io.WriteString(w, "Record succesfully deleteted")
	}
}

func (s *RESTServer) handleUpsert() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Record succesfully updated/created")
	}
}

func (s *RESTServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
