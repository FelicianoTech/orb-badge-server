package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-cli/api"
	"github.com/CircleCI-Public/circleci-cli/client"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Router            *mux.Router
	CertifiedOrbsOnly bool
}

type ShieldsJSON struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}

type Orb struct {
	namespace string
	name      string
	version   string
}

func (s *Server) initialize(certifiedOrbsOnly bool) {

	s.CertifiedOrbsOnly = certifiedOrbsOnly

	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/", s.viewHomepage).Methods("GET")
	s.Router.HandleFunc("/orb/{namespace}/{orb}", s.generateSchema).Methods("GET")
	// The path below is deprecated. Will be removed in a future release.
	s.Router.HandleFunc("/{namespace}/{orb}", s.generateSchema).Methods("GET")

	s.Router.Use(loggingMiddleware)
}

func (s *Server) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

func (s *Server) viewHomepage(w http.ResponseWriter, r *http.Request) {

	respondWithError(w, 403, "Go away.")
}

func (s *Server) generateSchema(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	orb := Orb{
		namespace: vars["namespace"],
		name:      vars["orb"],
	}

	gqlClient := client.NewClient("https://circleci.com", "graphql-unstable", "", false)

	orbVersion, err := api.OrbInfo(gqlClient, orb.namespace+"/"+orb.name)
	orb.version = orbVersion.Version

	if err != nil {
		fmt.Println("Failed to list Orbs.")
		fmt.Print(err)
	}

	shieldsJSON := ShieldsJSON{
		SchemaVersion: 1,
		Label:         "orb version",
		Message:       orb.version,
		Color:         "blue",
	}

	respondWithJSON(w, 200, shieldsJSON)
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

/*
 * Log requests automatically
 */
func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("[" + r.Method + "] " + r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
