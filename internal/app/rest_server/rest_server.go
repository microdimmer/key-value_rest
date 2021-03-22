package rest_server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/microdimmer/key-value_rest/internal/app/kv_db"
	"github.com/sirupsen/logrus"
)

type RESTServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	db     *kv_db.DataMap
}

type DataObject struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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

	s.logger.Info("server listening at port 8080")

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
	s.router.HandleFunc("/list", s.handleList()).Methods("GET")
	s.router.HandleFunc("/get/{key}", s.handleGet()).Methods("GET")
	s.router.HandleFunc("/delete/{key}", s.handleDelete()).Methods("DELETE")
	s.router.HandleFunc("/upsert", s.handleUpsert()).Methods("POST")
}

func (s *RESTServer) handleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, 200, s.db.List())
	}
}

func (s *RESTServer) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		s.logger.Info("get " + params["key"])
		record, ok := s.db.Get(params["key"])
		req := &DataObject{}
		if ok {
			req.Key = params["key"]
			req.Value = record
			respondWithJSON(w, 200, req)
		} else {
			respondWithJSON(w, 204, req)
		}
	}
}

func (s *RESTServer) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		s.logger.Info("delete " + params["key"])
		ok := s.db.Delete(params["key"])
		if ok {
			respondWithMessage(w, 200, "Record succesfully deleteted")
		} else {
			respondWithError(w, 400, "There is no record with key "+params["key"])
		}
	}
}

func (s *RESTServer) handleUpsert() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &DataObject{}
		err := json.NewDecoder(r.Body).Decode(req)
		s.logger.Info("upsert")
		if err != nil {
			respondWithError(w, 400, "Wrong payload")
		} else {
			s.db.Set(req.Key, req.Value)
			respondWithJSON(w, 200, req)
		}
	}
}

func respondWithMessage(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"message": message})
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
