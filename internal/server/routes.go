package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/health", s.healthHandler)

	r.HandlerFunc(http.MethodGet, "/getIngredientById/{id}", s.handleGetIngredientById)

	r.HandlerFunc(http.MethodGet, "/getIngredients", s.handleGetIngredients)

	r.HandlerFunc(http.MethodGet, "/getIngredientsByMenuId/{id}", s.handleGetIngredientsbyMenuId)

	return r
}

// TODO: Add Get Employee by Id

// TODO: Add Get Store by Id

// TODO: Add Get menu by Store Id

// TODO: Get Menu Items by Id

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) handleGetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := s.db.GetIngredients()
	if err != nil {
		log.Fatalf("error getting ingredients. Err: %v", err)
	}
	jsonResp, err := json.Marshal(ingredients)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}

func (s *Server) handleGetIngredientById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")
	ingredient, err := s.db.GetIngredientById(id)

	if err != nil {
		log.Fatalf("error getting ingredient by id. Err: %v", err)
	}

	jsonResp, err := json.Marshal(ingredient)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}

func (s *Server) handleGetIngredientsbyMenuId(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	ingredients, err := s.db.GetIngredientByMenuId(id)
	if err != nil {
		fmt.Println("error getting ingredients by menu id")
	}

	jsonResp, err := json.Marshal(ingredients)

	_, _ = w.Write(jsonResp)
}
